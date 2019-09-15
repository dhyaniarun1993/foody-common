package mongo

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is a wrapper on mongo.Collection with tracing Capability
type Collection struct {
	*mongo.Collection
	tracer opentracing.Tracer
}

// BulkWrite is a tracing wrapper around mongo collection BulkWrite
func (collection *Collection) BulkWrite(ctx context.Context, models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.BulkWrite",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.BulkWrite(ctx, models, opts...)
}

// InsertOne is a tracing wrapper around mongo collection InsertOne
func (collection *Collection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.InsertOne",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.InsertOne(ctx, document, opts...)
}

// InsertMany is a tracing wrapper around mongo collection InsertMany
func (collection *Collection) InsertMany(ctx context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.InsertMany",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.InsertMany(ctx, documents, opts...)
}

// DeleteOne is a tracing wrapper around mongo collection DeleteOne
func (collection *Collection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.DeleteOne",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.DeleteOne(ctx, filter, opts...)
}

// DeleteMany is a tracing wrapper around mongo collection DeleteMany
func (collection *Collection) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.DeleteMany",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.DeleteMany(ctx, filter, opts...)
}

// UpdateOne is a tracing wrapper around mongo collection UpdateOne
func (collection *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.UpdateOne",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany is a tracing wrapper around mongo collection UpdateMany
func (collection *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.UpdateMany",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.UpdateMany(ctx, filter, update, opts...)
}

// ReplaceOne is a tracing wrapper around mongo collection ReplaceOne
func (collection *Collection) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.ReplaceOne",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.ReplaceOne(ctx, filter, replacement, opts...)
}

// Aggregate is a tracing wrapper around mongo collection Aggregate
func (collection *Collection) Aggregate(ctx context.Context, pipeline interface{},
	opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.Aggregate",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.Aggregate(ctx, pipeline, opts...)
}

// CountDocuments is a tracing wrapper around mongo collection CountDocuments
func (collection *Collection) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions) (int64, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.CountDocuments",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.CountDocuments(ctx, filter, opts...)
}

// EstimatedDocumentCount is a tracing wrapper around mongo collection EstimatedDocumentCount
func (collection *Collection) EstimatedDocumentCount(ctx context.Context,
	opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.EstimatedDocumentCount",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.EstimatedDocumentCount(ctx, opts...)
}

// Distinct is a tracing wrapper around mongo collection Distinct
func (collection *Collection) Distinct(ctx context.Context, fieldName string, filter interface{},
	opts ...*options.DistinctOptions) ([]interface{}, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.Distinct",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.Distinct(ctx, fieldName, filter, opts...)
}

// Find is a tracing wrapper around mongo collection Find
func (collection *Collection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.Find",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.Find(ctx, filter, opts...)
}

// FindOne is a tracing wrapper around mongo collection FindOne
func (collection *Collection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.FindOne",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.FindOne(ctx, filter, opts...)
}

// FindOneAndDelete is a tracing wrapper around mongo collection FindOneAndDelete
func (collection *Collection) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.FindOneAndDelete",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.FindOneAndDelete(ctx, filter, opts...)
}

// FindOneAndReplace is a tracing wrapper around mongo collection FindOneAndReplace
func (collection *Collection) FindOneAndReplace(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.FindOneAndReplace",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.FindOneAndReplace(ctx, filter, replacement, opts...)
}

// FindOneAndUpdate is a tracing wrapper around mongo collection FindOneAndUpdate
func (collection *Collection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		collection.tracer.StartSpan(
			"Mongo.FindOneAndUpdate",
			opentracing.ChildOf(span.Context()),
		)
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	return collection.Collection.FindOneAndUpdate(ctx, filter, update, opts...)
}
