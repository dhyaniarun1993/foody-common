package mongo

import (
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is a wrapper on mongo.Database with tracing Capability
type Database struct {
	*mongo.Database
	tracer opentracing.Tracer
}

// Collection gets a handle for a given collection in the database.
func (db *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	collection := db.Database.Collection(name, opts...)
	return &Collection{collection, db.tracer}
}
