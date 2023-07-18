package conditions

import (
	"go-clean-architecture-example/pkg/db/mongo/utils"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The ConditionOperation is a simple type alias for the bson.M type and
// allows for more simple conversions
type ConditionOperation bson.M

// Condition Used in most condition operators. The Key is used to specify the field
// in the collection that is being queried and the value is the query itself.
// This query could be a new condition.Pipe
type Condition struct {
	Key   string
	Value interface{}
}

// Pipe The condition.Pipe differs from the aggregate.Pipe in that it doesn't
// produce a slice (bson.A) as it's output value, rather it produces a map
// (bson.M). This is because MongoDB queries use objects whereas the
// collection.aggregate function takes an array.
func Pipe(stage ...Condition) bson.M {
	d := bson.M{}
	for _, s := range stage {
		d[s.Key] = s.Value
	}
	return d
}

// cleanKey Helper function used internally to trim any whitespace on the Condition.Key
func cleanKey(key string) string {
	return strings.TrimSpace(key)
}

// buildBasicCondition Helper function used internally to construct basic conditions with the bson
// package. This essentially accepts an operation type, for example '$eq' or
// '$ne' as well as a field name (Key) and the value that's being queried on
// and turns it into it's JSON equivalent:
//
//	{
//		"Key": {
//			"operator": "Value"
//		}
//	}
func buildBasicCondition(operation string, conditions Condition) Condition {
	return Condition{
		Key: cleanKey(conditions.Key),
		Value: bson.M{
			operation: conditions.Value,
		},
	}
}

// InArray Uses the $in operator to check if a value exists in an array in a MongoDB
// document.
func InArray(c Condition) Condition {
	return buildBasicCondition("$in", c)
}

// ObjectIdMatch Uses the $eq operator and the util.StringToObjectId function to find a
// matching ObjectID in a MongoDB document.
func ObjectIdMatch(c Condition) Condition {
	c.Value = utils.StringToObjectId(c.Value.(string))
	return buildBasicCondition("$eq", c)
}

// BoolMatch Uses the $eq operator to filter on a matching bool value in a MongoDB document.
func BoolMatch(c Condition) Condition {
	return buildBasicCondition("$eq", c)
}

// NumberMatch Uses the $eq operator to filter on a matching number value in a MongoDB document.
func NumberMatch(c Condition) Condition {
	return buildBasicCondition("$eq", c)
}

// DateLessThanOrEqualTo Uses the $lte operator and the time.Parse function to filter documents where
// the document provided field is less than or equal to the specified field.
func DateLessThanOrEqualTo(c Condition) Condition {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", c.Value.(string))
	if err != nil {
		return Condition{}
	}
	c.Value = t
	return buildBasicCondition("$lte", c)
}

// DateGreaterThanOrEqualTo Uses the $gte operator and the time.Parse function to filter documents where
// the document provided field is greater than or equal to the specified field.
func DateGreaterThanOrEqualTo(c Condition) Condition {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", c.Value.(string))
	if err != nil {
		return Condition{}
	}
	c.Value = t
	return buildBasicCondition("$gte", c)
}

// EqualTo Uses the $eq operator to find a matching value of any type in a MongoDB
// document.
func EqualTo(c Condition) Condition {
	return buildBasicCondition("$eq", c)
}

// NotEqualTo Uses the $ne operator to filter values that do not match the provided value
// of any type in a MongoDB document.
func NotEqualTo(c Condition) Condition {
	return buildBasicCondition("$ne", Condition{
		Key:   c.Key,
		Value: c.Value,
	})
}

// ElemMatch Uses the $elemMatch operator to matche documents that contain an array field
// with at least one element that matches all the specified query criteria.
func ElemMatch(c Condition) Condition {
	return buildBasicCondition("$elemMatch", c)
}

// StringMatch Uses the $regex operator to match a whole string in a MongoDB document.
func StringMatch(c Condition) Condition {
	return buildBasicCondition(
		"$regex",
		Condition{
			Key: c.Key,
			Value: primitive.Regex{
				Pattern: "^" + c.Value.(string) + "$",
				Options: "i",
			},
		})
}

// StringLike Uses the $regex operator to match any part of a string in a MongoDB document.
func StringLike(c Condition) Condition {
	return buildBasicCondition(
		"$regex",
		Condition{
			Key: c.Key,
			Value: primitive.Regex{
				Pattern: c.Value.(string),
				Options: "i",
			},
		})
}

// StringStartsWith Uses the $regex operator to match a string that starts with the provided
// string in a MongoDB document.
func StringStartsWith(c Condition) Condition {
	return buildBasicCondition(
		"$regex",
		Condition{
			Key: c.Key,
			Value: primitive.Regex{
				Pattern: "^" + c.Value.(string) + "",
				Options: "i",
			},
		})
}

// StringEndsWith Uses the $regex operator to match a string that ends with the provided
// string in a MongoDB document.
func StringEndsWith(c Condition) Condition {
	return buildBasicCondition(
		"$regex",
		Condition{
			Key: c.Key,
			Value: primitive.Regex{
				Pattern: c.Value.(string) + "$",
				Options: "i",
			},
		})
}
