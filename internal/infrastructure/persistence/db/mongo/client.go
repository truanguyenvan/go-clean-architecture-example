package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Client struct {
	database       *mongo.Database
	collectionName string
}

func NewClient(client *mongo.Client, dbName, collection string) *Client {
	return &Client{
		database:       client.Database(dbName),
		collectionName: collection,
	}
}

func (client *Client) WithCollection(collectionName string) *Client {
	client.collectionName = collectionName
	return client
}
