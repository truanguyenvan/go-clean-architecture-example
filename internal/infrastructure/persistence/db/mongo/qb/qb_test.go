package qb

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

//go:embed example_query.json
var query string

//go:embed example_aggregate.json
var aggregate string

func TestBuild(t *testing.T) {
	// Build the query
	out, err := Build[bson.M](query, map[string]any{
		"Foo":  "000000000000000000000000",
		"Foo2": hex("000000000000000000000000"),
		"Expr": "/john|jane/i",
	})
	assert.NoError(t, err)

	// Assert the output
	assert.Equal(t, bson.M{"$match": bson.M{
		"foo":  hex("000000000000000000000000"),
		"foo2": hex("000000000000000000000000"),
		"bar": bson.M{"$regex": bson.M{
			"Pattern": "john|jane",
			"Options": "i",
		}},
	}}, out)
}

func TestAggregate(t *testing.T) {
	out, err := Build[bson.A](aggregate, map[string]any{
		"Foo": "000000000000000000000000",
	})
	assert.NoError(t, err)

	assert.Equal(t, bson.A{bson.D{primitive.E{
		Key: "$match",
		Value: primitive.D{
			primitive.E{Key: "foo", Value: hex("000000000000000000000000")},
		},
	}}}, out)
}

func hex(h string) primitive.ObjectID {
	k, err := primitive.ObjectIDFromHex(h)
	if err != nil {
		panic(err)
	}
	return k
}
