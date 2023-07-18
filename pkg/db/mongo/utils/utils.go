package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// OptionalString Normally when unmarshalling bson into a struct, all values that are missing
// from the database are initialized as their zero value (with the exception of
// pointers whose zero value is nil). If a value is specified in the struct but
// is missing in the database it will throw an error. The OptionalString type
// allows you specify strings that can be considered optional. Fields that are
// strings and are not provided by the database are initialized as their zero
// value.
type OptionalString string

func (o *OptionalString) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	if t == bsontype.Null {
		*o = ""
		return nil
	}

	rv := bson.RawValue{
		Type:  t,
		Value: data,
	}

	err := rv.Unmarshal((*string)(o))
	if err != nil {
		return fmt.Errorf("unmarshaling optional string: %w\n", err)
	}
	return nil
}

// Helper function for determining if an optional string is set to it's zero value.
func (o *OptionalString) IsZero() bool {
	return len(*o) == 0
}

// StringToObjectId Converts a string to a primitive.ObjectID. If there is an error in the conversion
// it will return a new ObjectID generated from the current time.
func StringToObjectId(id string) primitive.ObjectID {
	o, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NewObjectID()
	}
	return o
}

// StringsToObjectId Converts a slice of strings to a slice of primitive.ObjectID. If there is
// an error in the conversion of any one of the strings, it will return a new
// ObjectID generated from the current time.
func StringsToObjectId(ids []string) []primitive.ObjectID {
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		o, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			o = primitive.NewObjectID()
		}
		objectIds = append(objectIds, o)
	}
	return objectIds
}

// PrintJson Helper function for debugging pipelines by decoding them into json. You
// can take the JSON use it to query the database directly for testing purposes.
func PrintJson(d interface{}) {
	data, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatalf("marshaling json: %v", err)
	}
	fmt.Printf("%v\n", string(data))
}

// PtrBool Returns a pointer to a provided bool value. Used mainly by internal gomongo
// packages.
func PtrBool(val bool) *bool {
	v := val
	return &v
}

// PtrString Returns a pointer to a provided string value. Used mainly by internal gomongo
// packages.
func PtrString(val string) *string {
	v := val
	return &v
}

// PtrInt Returns a pointer to a provided int value. Used mainly by internal gomongo
// packages.
func PtrInt(val int) *int {
	v := val
	return &v
}

// PtrTimeNow Returns a pointer to a the current time value. Used mainly by internal gomongo
// packages.
func PtrTimeNow() *time.Time {
	n := time.Now()
	return &n
}

// PtrObjectIdSlice Returns a pointer to a slice from a provided primtive.ObjectID. Used mainly
// by internal gomongo packages.
func PtrObjectIdSlice(objectId primitive.ObjectID) *[]primitive.ObjectID {
	ns := &[]primitive.ObjectID{objectId}
	return ns
}
