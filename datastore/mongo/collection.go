package mongo

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.BulkWrite", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("models::%v", models))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.BulkWrite(ctx, models, opts...)
}

// InsertOne is a tracing wrapper around mongo collection InsertOne
func (collection *Collection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.InsertOne", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("document::%v", document))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.InsertOne(ctx, document, opts...)
}

// InsertMany is a tracing wrapper around mongo collection InsertMany
func (collection *Collection) InsertMany(ctx context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.InsertMany", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("documents::%v", documents))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.InsertMany(ctx, documents, opts...)
}

// DeleteOne is a tracing wrapper around mongo collection DeleteOne
func (collection *Collection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.DeleteOne", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.DeleteOne(ctx, filter, opts...)
}

// DeleteMany is a tracing wrapper around mongo collection DeleteMany
func (collection *Collection) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.DeleteMany", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.DeleteMany(ctx, filter, opts...)
}

// UpdateOne is a tracing wrapper around mongo collection UpdateOne
func (collection *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.UpdateOne", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v, update::%v", filter, update))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany is a tracing wrapper around mongo collection UpdateMany
func (collection *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.UpdateMany", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v, update::%v", filter, update))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.UpdateMany(ctx, filter, update, opts...)
}

// ReplaceOne is a tracing wrapper around mongo collection ReplaceOne
func (collection *Collection) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.ReplaceOne", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v, replacement::%v", filter, replacement))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.ReplaceOne(ctx, filter, replacement, opts...)
}

// Aggregate is a tracing wrapper around mongo collection Aggregate
func (collection *Collection) Aggregate(ctx context.Context, pipeline interface{},
	opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.Aggregate", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("pipeline::%v", pipeline))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.Aggregate(ctx, pipeline, opts...)
}

// CountDocuments is a tracing wrapper around mongo collection CountDocuments
func (collection *Collection) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions) (int64, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.CountDocuments", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.CountDocuments(ctx, filter, opts...)
}

// EstimatedDocumentCount is a tracing wrapper around mongo collection EstimatedDocumentCount
func (collection *Collection) EstimatedDocumentCount(ctx context.Context,
	opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.EstimatedDocumentCount", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.EstimatedDocumentCount(ctx, opts...)
}

// Distinct is a tracing wrapper around mongo collection Distinct
func (collection *Collection) Distinct(ctx context.Context, fieldName string, filter interface{},
	opts ...*options.DistinctOptions) ([]interface{}, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.Distinct", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("fieldName::%v, filter::%v", fieldName, filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.Distinct(ctx, fieldName, filter, opts...)
}

// Find is a tracing wrapper around mongo collection Find
func (collection *Collection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.Find", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.Find(ctx, filter, opts...)
}

// FindOne is a tracing wrapper around mongo collection FindOne
func (collection *Collection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.FindOne", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.FindOne(ctx, filter, opts...)
}

// FindOneAndDelete is a tracing wrapper around mongo collection FindOneAndDelete
func (collection *Collection) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.FindOneAndDelete", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v", filter))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.FindOneAndDelete(ctx, filter, opts...)
}

// FindOneAndReplace is a tracing wrapper around mongo collection FindOneAndReplace
func (collection *Collection) FindOneAndReplace(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.FindOneAndReplace", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v, replacement::%v", filter, replacement))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.FindOneAndReplace(ctx, filter, replacement, opts...)
}

// FindOneAndUpdate is a tracing wrapper around mongo collection FindOneAndUpdate
func (collection *Collection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := collection.tracer.StartSpan(
			fmt.Sprintf("Mongo.%v.FindOneAndUpdate", collection.Name()),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "mongo.Client")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, fmt.Sprintf("filter::%v, replacement::%v", filter, update))
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return collection.Collection.FindOneAndUpdate(ctx, filter, update, opts...)
}
