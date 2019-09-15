package mongo

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Configuration provides configuration for MongoDB driver
type Configuration struct {
	URI      string `required:"true" split_words:"true"`
	Database string `required:"true"`
}

// Client is a wrapper on mongo.Client with tracing Capability
type Client struct {
	*mongo.Client
	tracer opentracing.Tracer
}

// CreateMongoDBPool creates connection pool for MongoDB server
func CreateMongoDBPool(configuration Configuration, tracer opentracing.Tracer) *Client {
	connectCtx, connectCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer connectCancel()
	clientOptions := options.Client().ApplyURI(configuration.URI)
	client, connectError := mongo.Connect(connectCtx, clientOptions)
	if connectError != nil {
		panic(connectError)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancel()
	pingError := client.Ping(pingCtx, readpref.Primary())
	if pingError != nil {
		panic(pingError)
	}

	return &Client{client, tracer}
}

// Database returns a handle for a given database.
func (client *Client) Database(name string, opts ...*options.DatabaseOptions) *Database {
	database := client.Client.Database(name, opts...)
	return &Database{database, client.tracer}
}
