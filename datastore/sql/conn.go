package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Conn is a wrapper on sql.Conn with tracing Capability
type Conn struct {
	*sql.Conn
	tracer opentracing.Tracer
}

// BeginTx instruments the sql.DB BeginTx with tracing capability
func (conn *Conn) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("conn.beginTxn"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return conn.Conn.BeginTx(ctx, opts)
}

// ExecContext instruments the sql.DB ExecContext with tracing capability
func (conn *Conn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("conn.execute"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return conn.Conn.ExecContext(ctx, query, args...)
}

// PingContext instruments the sql.DB PingContext with tracing capability
func (conn *Conn) PingContext(ctx context.Context) error {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("conn.ping"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return conn.Conn.PingContext(ctx)
}

// PrepareContext instruments the sql.DB PrepareContext with tracing capability
func (conn *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("db.prepare"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	stmt, err := conn.Conn.PrepareContext(ctx, query)
	return &Stmt{stmt, conn.tracer}, err
}

// QueryContext instruments the sql.DB QueryContext with tracing capability
func (conn *Conn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("db.query"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return conn.Conn.QueryContext(ctx, query, args...)
}

// QueryRowContext instruments the sql.DB QueryRowContext with tracing capability
func (conn *Conn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := conn.tracer.StartSpan(
			fmt.Sprintf("db.queryRow"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return conn.Conn.QueryRowContext(ctx, query, args...)
}
