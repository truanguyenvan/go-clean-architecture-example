package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Config struct {
	URI               string
	UserName          string
	Password          string
	ConnectionTimeout time.Duration
	MaxConnIdleTime   time.Duration
	MinPoolSize       uint64
	MaxPoolSize       uint64
	DbName            string
	Collection        string
}

type Mongo struct {
	_client *mongo.Client
	Client  *Client
	ctx     context.Context
}

// Open - creates a new Mongo
func Open(cfg Config) (*Mongo, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(cfg.URI).
			SetAuth(options.Credential{Username: cfg.UserName, Password: cfg.Password}).
			SetConnectTimeout(cfg.ConnectionTimeout).
			SetMaxConnIdleTime(cfg.MaxConnIdleTime).
			SetMinPoolSize(cfg.MinPoolSize).
			SetMaxPoolSize(cfg.MaxPoolSize))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Mongo{
		ctx:     ctx,
		_client: client,
		Client:  NewClient(client, cfg.DbName, cfg.Collection),
	}, nil
}

// Disconnect - used mainly in testing to avoid capping out the concurrent connections on MongoDB
func (m *Mongo) Disconnect() {
	err := m._client.Disconnect(m.ctx)
	if err != nil {
		log.Fatalf("disconnecting from mongodb: %v", err)
	}
}

// Ping sends a ping command to verify that the client can connect to the deployment.
func (m *Mongo) Ping() error {
	return m._client.Ping(m.ctx, nil)
}

// WithDatabase - create a new Client
func (m *Mongo) WithDatabase(dbName, collection string) *Client {
	return NewClient(m._client, dbName, collection)
}
