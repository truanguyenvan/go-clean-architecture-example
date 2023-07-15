package mongo

import (
	"context"
	"fmt"
	"go-clean-architecture-example/internal/infrastructure/persistence/db/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find returns all documents that match the provided query
func (client *Client) Find(ctx context.Context, query interface{}, opts ...*options.FindOptions) (rtn interface{}, err error) {
	c, err := client.database.Collection(client.collectionName).Find(ctx, query, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Find: %v", err)
	}

	err = c.All(ctx, &rtn)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Find - parsing cursor: %v", err)
	}
	return
}

// FindCursor returns all documents that match the provided query and return the cursor
func (client *Client) FindCursor(ctx context.Context, query interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	c, err := client.database.Collection(client.collectionName).Find(ctx, query, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Find: %v", err)
	}
	return c, nil
}

// FindOne returns the first document that matches the provided query
func (client *Client) FindOne(ctx context.Context, query interface{}, opts ...*options.FindOneOptions) (rtn interface{}, err error) {
	err = client.database.Collection(client.collectionName).FindOne(ctx, query, opts...).Decode(&rtn)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		return nil, fmt.Errorf("Mongo.FindOne: %v", err)
	}
	return
}

// FindById returns the first document that has an _id that matches the id parameter
func (client *Client) FindById(ctx context.Context, id string) (rtn interface{}, err error) {
	query := bson.M{"_id": bson.M{"$eq": utils.StringToObjectId(id)}}
	rtn, err = client.FindOne(ctx, query)
	return
}

// FindByIdList returns all documents where the _id is matches an id specified in the ids []string parameter
func (client *Client) FindByIdList(ctx context.Context, ids []string) (rtn interface{}, err error) {
	query := bson.M{"_id": bson.M{"$in": utils.StringsToObjectId(ids)}}
	rtn, err = client.Find(ctx, query)
	return
}

func (client *Client) FindDistinct(ctx context.Context, fieldName string, query interface{}) (rtn []interface{}, err error) {
	if fieldName == "" {
		return nil, fmt.Errorf("Client.FindDistinct - requires a fieldName")
	}

	rtn, err = client.database.Collection(client.collectionName).Distinct(ctx, fieldName, query)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Distinct: %v", err)
	}
	return
}

func (client *Client) Aggregate(ctx context.Context, pipeline interface{}, allowDiskUse bool) (*mongo.Cursor, error) {
	opts := options.AggregateOptions{
		AllowDiskUse: utils.PtrBool(allowDiskUse),
	}

	c, err := client.database.Collection(client.collectionName).Aggregate(ctx, pipeline, &opts)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Aggregate: %v", err)
	}
	return c, nil
}

func (client *Client) Create(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	op, err := client.database.Collection(client.collectionName).InsertOne(ctx, doc, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.InsertOne: %v", err)
	}
	return op, nil
}

func (client *Client) CreateMany(ctx context.Context, docs []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	op, err := client.database.Collection(client.collectionName).InsertMany(ctx, docs, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.InsertMany: %v", err)
	}
	return op, nil
}

func (client *Client) UpdateOne(ctx context.Context, query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	op, err := client.database.Collection(client.collectionName).UpdateOne(ctx, query, update, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.UpdateOne: %v", err)
	}
	return op, nil
}

func (client *Client) UpdateMany(ctx context.Context, query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	op, err := client.database.Collection(client.collectionName).UpdateMany(ctx, query, update, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.UpdateMany: %v", err)
	}
	return op, nil
}

// UpdateById Updates a single document whose _id matches what is provided in the `id` parameter
func (client *Client) UpdateById(ctx context.Context, id string, update interface{}, replace bool) (*mongo.UpdateResult, error) {
	q := bson.M{
		"_id": bson.M{
			"$eq": utils.StringToObjectId(id),
		},
	}

	if !replace {
		update = bson.M{
			"$set": update,
		}
	}

	result, err := client.UpdateOne(ctx, q, update)
	return result, err
}

// UpdateByIdList Updates a set of documents whose _ids match what is provided in the `ids` parameter
func (client *Client) UpdateByIdList(ctx context.Context, ids []string, update interface{}, replace bool) (*mongo.UpdateResult, error) {
	q := bson.M{
		"_id": bson.M{
			"$in": utils.StringsToObjectId(ids),
		},
	}

	if !replace {
		update = bson.M{
			"$set": update,
		}
	}

	result, err := client.UpdateMany(ctx, q, update)

	return result, err
}

func (client *Client) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	op, err := client.database.Collection(client.collectionName).DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.DeleteOne: %v", err)
	}
	return op, err
}

func (client *Client) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	op, err := client.database.Collection(client.collectionName).DeleteMany(ctx, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.DeleteMany: %v", err)
	}
	return op, err
}

// DeleteById Deletes a document with a corresponding _id
func (client *Client) DeleteById(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": utils.StringToObjectId(id),
		},
	}

	op, err := client.DeleteOne(ctx, filter)
	return op, err
}

// QuickDeleteById Works similar to DeleteById but doesn't return an op result or an error
func (client *Client) QuickDeleteById(ctx context.Context, id string) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": utils.StringToObjectId(id),
		},
	}
	_, _ = client.DeleteOne(ctx, filter)
}

// DeleteByIdList Updates a set of documents whose _ids match what is provided in the `ids` parameter
func (client *Client) DeleteByIdList(ctx context.Context, ids []string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"_id": bson.M{
			"$in": utils.StringsToObjectId(ids),
		},
	}

	op, err := client.DeleteOne(ctx, filter)
	return op, err
}

func (client *Client) Watch(ctx context.Context) (*mongo.ChangeStream, error) {
	cs, err := client.database.Collection(client.collectionName).Watch(ctx, mongo.Pipeline{})
	if err != nil {
		return nil, fmt.Errorf("Mongo.Watch: %v", err)
	}
	return cs, nil
}
