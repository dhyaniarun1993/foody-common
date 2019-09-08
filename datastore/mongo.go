package datastore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Configuration provides configuration for MongoDB driver
type Configuration struct {
	URI      string `required:"true" split_words:"true"`
	Database string `required:"true"`
}

// CreateMongoDBPool creates connection pool for MongoDB server
func CreateMongoDBPool(configuration Configuration) *mongo.Client {
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

	return client
}
