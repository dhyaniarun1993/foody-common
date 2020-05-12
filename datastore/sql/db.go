package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// DB is a wrapper on sql.DB with tracing Capability
type DB struct {
	*sql.DB
	tracer opentracing.Tracer
}

// BeginTx instruments the sql.DB BeginTx with tracing capability
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.beginTxn"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return db.DB.BeginTx(ctx, opts)
}

// Conn instruments the sql.DB Conn with tracing capability
func (db *DB) Conn(ctx context.Context) (*Conn, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.conn"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	conn, err := db.DB.Conn(ctx)
	return &Conn{conn, db.tracer}, err
}

// ExecContext instruments the sql.DB ExecContext with tracing capability
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.execute"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return db.DB.ExecContext(ctx, query, args...)
}

// PingContext instruments the sql.DB PingContext with tracing capability
func (db *DB) PingContext(ctx context.Context) error {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.ping"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return db.DB.PingContext(ctx)
}

// PrepareContext instruments the sql.DB PrepareContext with tracing capability
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.prepare"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	stmt, err := db.DB.PrepareContext(ctx, query)
	return &Stmt{stmt, db.tracer}, err
}

// QueryContext instruments the sql.DB QueryContext with tracing capability
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.query"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return db.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext instruments the sql.DB QueryRowContext with tracing capability
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := db.tracer.StartSpan(
			fmt.Sprintf("db.queryRow"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return db.DB.QueryRowContext(ctx, query, args...)
}
