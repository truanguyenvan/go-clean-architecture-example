package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Nexters/myply/infrastructure/configs"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database // connection
}

func NewMongoDB(config *configs.Config) (*MongoInstance, error) {
	mongoURI := config.MongoURI + "/" + config.MongoDBName
	clientOptions := options.Client().ApplyURI(mongoURI)

	switch config.Phase {
	case configs.Production:
		clientOptions.SetHeartbeatInterval(15 * time.Second)
		clientOptions.SetMaxPoolSize(100)
		clientOptions.SetMinPoolSize(1)
		clientOptions.SetMaxConnIdleTime(10 * time.Second)
	default:
		clientOptions.SetHeartbeatInterval(10 * time.Second) // default 10s
		clientOptions.SetMaxPoolSize(100)                    // default 100
		clientOptions.SetMinPoolSize(0)                      // default 0
		clientOptions.SetMaxConnIdleTime(0)                  // default 0s
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {

		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.MongoTTL)

	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(config.MongoDBName)

	if err != nil {
		return nil, err
	}

	return &MongoInstance{
		Client: client,
		Db:     db,
	}, nil
}
