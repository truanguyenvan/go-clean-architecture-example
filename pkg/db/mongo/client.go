package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client interface {
	WithCollection(collectionName string) Client
	WithContext(config ContextConfig) Client
	Find(query interface{}, opts ...*options.FindOptions) (rtn interface{}, err error)
	FindCursor(query interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(query interface{}, opts ...*options.FindOneOptions) (rtn interface{}, err error)
	FindById(id string) (rtn interface{}, err error)
	FindByIdList(ids []string) (rtn interface{}, err error)
	FindDistinct(fieldName string, query interface{}) (rtn []interface{}, err error)
	Aggregate(pipeline interface{}, allowDiskUse bool) (*mongo.Cursor, error)
	Create(doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	CreateMany(docs []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(query interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateById(id string, update interface{}, replace bool) (*mongo.UpdateResult, error)
	UpdateByIdList(ids []string, update interface{}, replace bool) (*mongo.UpdateResult, error)
	DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteById(id string) (*mongo.DeleteResult, error)
	QuickDeleteById(id string)
	DeleteByIdList(ids []string) (*mongo.DeleteResult, error)
	Watch(ctx context.Context) (*mongo.ChangeStream, error)
}

type ContextConfig struct {
	ctx        context.Context
	ctxTimeout time.Duration
}

type ClientConfig struct {
	ContextConfig
	dbName         string
	collectionName string
}

type client struct {
	ClientConfig
	database *mongo.Database
}

func newClient(mgoClient *mongo.Client, config ClientConfig) Client {
	return &client{
		ClientConfig: config,
		database:     mgoClient.Database(config.dbName),
	}
}

func (clt client) WithCollection(collectionName string) Client {
	clt.collectionName = collectionName
	return &clt
}

func (clt client) WithContext(config ContextConfig) Client {
	clt.ContextConfig = config
	return &clt
}
