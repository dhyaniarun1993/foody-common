package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Stmt is a wrapper on sql.Stmt with tracing Capability
type Stmt struct {
	*sql.Stmt
	tracer opentracing.Tracer
}

// ExecContext instruments the sql.Tx ExecContext with tracing capability
func (stmt *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := stmt.tracer.StartSpan(
			fmt.Sprintf("stmt.execute"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return stmt.Stmt.ExecContext(ctx, args...)
}

// QueryContext instruments the sql.Tx QueryContext with tracing capability
func (stmt *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := stmt.tracer.StartSpan(
			fmt.Sprintf("stmt.query"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return stmt.Stmt.QueryContext(ctx, args...)
}

// QueryRowContext instruments the sql.Tx QueryRowContext with tracing capability
func (stmt *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		newSpan := stmt.tracer.StartSpan(
			fmt.Sprintf("stmt.queryRow"),
			opentracing.ChildOf(span.Context()),
		)
		ext.Component.Set(newSpan, "database/sql")
		ext.SpanKind.Set(newSpan, "client")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
		defer newSpan.Finish()
	}
	return stmt.Stmt.QueryRowContext(ctx, args...)
}
