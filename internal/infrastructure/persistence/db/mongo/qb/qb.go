package qb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
	"text/template"
)

// Builtins contains all the currently registered functions that can be used when
// templating query templates. Publicly exposed so that callers can add their own
// conversion functions.
var Builtins = map[string]interface{}{
	"oid":   oid,
	"regex": regex,
}

// Build builds a MongoDB query of type T given an input query template. The caller may optionally
// provide variables that will be templated into the query. The caller may also use any of the registered
// functions above to handle special types of variables such as the ObjectIds or regular expressions.
func Build[T bson.M | bson.D | bson.A | bson.E](query string, vars map[string]any) (out T, err error) {
	var buf bytes.Buffer
	tmpl, err := template.New("").Funcs(Builtins).Parse(query)
	if err != nil {
		return out, fmt.Errorf("parsing template: %w", err)
	}
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return out, fmt.Errorf("executing template: %w", err)
	}
	return out, bson.UnmarshalExtJSON(buf.Bytes(), false, &out)
}

// oid a string or an ObjectId input parameter to a $oid in the MongoDB query.
func oid(value reflect.Value) (interface{}, error) {
	str, err := wrap("$oid", value.Interface())
	if err != nil {
		return nil, err
	}
	return str, nil
}

// regex accepts a single parameter containing a PCRE regular expression formatted
// like the example in the MongoDB documentation (https://www.mongodb.com/docs/manual/reference/operator/query/regex/).
// where the string looks like "/<expr>/<options>", i.e. "/john/i".
func regex(value reflect.Value) (interface{}, error) {
	v, ok := value.Interface().(string)
	if !ok {
		return nil, fmt.Errorf("must provide a string, got %T", value.Interface())
	}
	if len(v) == 0 {
		return nil, fmt.Errorf("must provide a valid regex")
	}
	if !strings.HasPrefix(v, "/") {
		return nil, fmt.Errorf("regex must start with '/' character")
	}
	parts := strings.Split(v, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("expected regex to match format '/<expr>/<options>'")
	}
	str, err := wrap("$regex", primitive.Regex{
		Pattern: parts[1],
		Options: parts[2],
	})
	if err != nil {
		return nil, err
	}
	return str, nil
}

// wrap wraps a value in a key. It is used to wrap values with custom MongoDB operators
// like the $oid or $regex operators.
func wrap(key string, value any) (string, error) {
	b, err := json.Marshal(map[string]any{key: value})
	if err != nil {
		return "", err
	}
	return string(b), nil
}
