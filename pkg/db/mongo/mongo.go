package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
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
	ContextTimeout    time.Duration
}

type Mongo struct {
	_client *mongo.Client
	Client  Client
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
			SetMaxPoolSize(cfg.MaxPoolSize).
			SetReadPreference(readpref.SecondaryPreferred()))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	if cfg.ContextTimeout == 0 {
		cfg.ConnectionTimeout = 30 * time.Second
	}
	clientConfig := ClientConfig{
		ContextConfig: ContextConfig{
			ctx:        ctx,
			ctxTimeout: cfg.ContextTimeout,
		},
		dbName:         cfg.DbName,
		collectionName: cfg.Collection,
	}

	return &Mongo{
		ctx:     ctx,
		_client: new(mongo.Client),
		Client:  newClient(new(mongo.Client), clientConfig),
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
	return m._client.Ping(m.ctx, readpref.Primary())
}

// WithDatabase - create a new Client
func (m *Mongo) WithDatabase(clientConfig ClientConfig) Client {
	return newClient(m._client, clientConfig)
}

func (m *Mongo) Transaction(txnFn func(sc mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	session, err := m._client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(m.ctx)

	result, err := session.WithTransaction(m.ctx, txnFn,
		options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.Majority()))
	if err != nil {
		return nil, err
	}
	return result, nil
}
