package query

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Note: DBTX and *Queries generated by sqlc

type tracingDBTX struct {
	db DBTX
}

func NewWithTracing(conn DBTX) *Queries {
	return New(&tracingDBTX{db: conn})
}

func (t *tracingDBTX) Exec(ctx context.Context, s string, i ...interface{}) (pgconn.CommandTag, error) {
	ctx, span := createSpan(ctx, s)
	defer span.End()
	return t.db.Exec(ctx, s, i...)
}

func (t *tracingDBTX) Query(ctx context.Context, s string, i ...interface{}) (pgx.Rows, error) {
	ctx, span := createSpan(ctx, s)
	defer span.End()
	return t.db.Query(ctx, s, i...)
}

func (t *tracingDBTX) QueryRow(ctx context.Context, s string, i ...interface{}) pgx.Row {
	ctx, span := createSpan(ctx, s)
	defer span.End()
	return t.db.QueryRow(ctx, s, i...)
}

func (t *tracingDBTX) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	ctx, span := createSpan(ctx, "copyFrom:"+tableName.Sanitize())
	defer span.End()
	return t.db.CopyFrom(ctx, tableName, columnNames, rowSrc)
}