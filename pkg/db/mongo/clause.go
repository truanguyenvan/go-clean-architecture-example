package mongo

import (
	"context"
	"fmt"
	"go-clean-architecture-example/pkg/db/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (clt client) getTemporaryCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(clt.ctx, clt.ctxTimeout)
}

// Find returns all documents that match the provided query
func (clt client) Find(query interface{}, opts ...*options.FindOptions) (rtn interface{}, err error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	c, err := clt.database.Collection(clt.collectionName).Find(ctx, query, opts...)
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
func (clt client) FindCursor(query interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	c, err := clt.database.Collection(clt.collectionName).Find(ctx, query, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Find: %v", err)
	}
	return c, nil
}

// FindOne returns the first document that matches the provided query
func (clt client) FindOne(query interface{}, opts ...*options.FindOneOptions) (rtn interface{}, err error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	err = clt.database.Collection(clt.collectionName).FindOne(ctx, query, opts...).Decode(&rtn)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		return nil, fmt.Errorf("Mongo.FindOne: %v", err)
	}
	return
}

// FindById returns the first document that has an _id that matches the id parameter
func (clt client) FindById(id string) (rtn interface{}, err error) {
	query := bson.M{"_id": bson.M{"$eq": utils.StringToObjectId(id)}}
	rtn, err = clt.FindOne(query)
	return
}

// FindByIdList returns all documents where the _id is matches an id specified in the ids []string parameter
func (clt client) FindByIdList(ids []string) (rtn interface{}, err error) {
	query := bson.M{"_id": bson.M{"$in": utils.StringsToObjectId(ids)}}
	rtn, err = clt.Find(query)
	return
}

func (clt client) FindDistinct(fieldName string, query interface{}) (rtn []interface{}, err error) {
	if fieldName == "" {
		return nil, fmt.Errorf("Client.FindDistinct - requires a fieldName")
	}

	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	rtn, err = clt.database.Collection(clt.collectionName).Distinct(ctx, fieldName, query)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Distinct: %v", err)
	}
	return
}

func (clt client) Aggregate(pipeline interface{}, allowDiskUse bool) (*mongo.Cursor, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	opts := options.AggregateOptions{
		AllowDiskUse: utils.PtrBool(allowDiskUse),
	}
	c, err := clt.database.Collection(clt.collectionName).Aggregate(ctx, pipeline, &opts)
	if err != nil {
		return nil, fmt.Errorf("Mongo.Aggregate: %v", err)
	}
	return c, nil
}

func (clt client) Create(doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).InsertOne(ctx, doc, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.InsertOne: %v", err)
	}
	return op, nil
}

func (clt client) CreateMany(docs []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).InsertMany(ctx, docs, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.InsertMany: %v", err)
	}
	return op, nil
}

func (clt client) UpdateOne(query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).UpdateOne(ctx, query, update, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.UpdateOne: %v", err)
	}
	return op, nil
}

func (clt client) UpdateMany(query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).UpdateMany(ctx, query, update, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.UpdateMany: %v", err)
	}
	return op, nil
}

// UpdateById Updates a single document whose _id matches what is provided in the `id` parameter
func (clt client) UpdateById(id string, update interface{}, replace bool) (*mongo.UpdateResult, error) {
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

	result, err := clt.UpdateOne(q, update)
	return result, err
}

// UpdateByIdList Updates a set of documents whose _ids match what is provided in the `ids` parameter
func (clt client) UpdateByIdList(ids []string, update interface{}, replace bool) (*mongo.UpdateResult, error) {
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

	result, err := clt.UpdateMany(q, update)
	return result, err
}

func (clt client) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.DeleteOne: %v", err)
	}
	return op, err
}

func (clt client) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	op, err := clt.database.Collection(clt.collectionName).DeleteMany(ctx, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("Mongo.DeleteMany: %v", err)
	}
	return op, err
}

// DeleteById Deletes a document with a corresponding _id
func (clt client) DeleteById(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": utils.StringToObjectId(id),
		},
	}

	op, err := clt.DeleteOne(filter)
	return op, err
}

// QuickDeleteById Works similar to DeleteById but doesn't return an op result or an error
func (clt client) QuickDeleteById(id string) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": utils.StringToObjectId(id),
		},
	}
	_, _ = clt.DeleteOne(filter)
}

// DeleteByIdList Updates a set of documents whose _ids match what is provided in the `ids` parameter
func (clt client) DeleteByIdList(ids []string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"_id": bson.M{
			"$in": utils.StringsToObjectId(ids),
		},
	}

	op, err := clt.DeleteOne(filter)
	return op, err
}

func (clt client) Watch(ctx context.Context) (*mongo.ChangeStream, error) {
	ctx, cancel := clt.getTemporaryCtx()
	defer cancel()

	cs, err := clt.database.Collection(clt.collectionName).Watch(ctx, mongo.Pipeline{})
	if err != nil {
		return nil, fmt.Errorf("Mongo.Watch: %v", err)
	}
	return cs, nil
}
