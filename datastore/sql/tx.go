package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Tx is a wrapper on sql.Tx with tracing Capability
type Tx struct {
	*sql.Tx
	tracer opentracing.Tracer
}

// ExecContext instruments the sql.Tx ExecContext with tracing capability
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := tx.tracer.StartSpan(
			fmt.Sprintf("txn.execute"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return tx.Tx.ExecContext(ctx, query, args...)
}

// PrepareContext instruments the sql.Tx PrepareContext with tracing capability
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := tx.tracer.StartSpan(
			fmt.Sprintf("txn.prepare"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	stmt, err := tx.Tx.PrepareContext(ctx, query)
	return &Stmt{stmt, tx.tracer}, err
}

// QueryContext instruments the sql.Tx QueryContext with tracing capability
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := tx.tracer.StartSpan(
			fmt.Sprintf("txn.query"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return tx.Tx.QueryContext(ctx, query, args...)
}

// QueryRowContext instruments the sql.Tx QueryRowContext with tracing capability
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := tx.tracer.StartSpan(
			fmt.Sprintf("txn.queryRow"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ext.DBStatement.Set(newSpan, query)
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return tx.Tx.QueryRowContext(ctx, query, args...)
}

// StmtContext instruments the sql.Tx StmtContext with tracing capability
func (tx *Tx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := tx.tracer.StartSpan(
			fmt.Sprintf("txn.statement"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return tx.Tx.StmtContext(ctx, stmt)
}
