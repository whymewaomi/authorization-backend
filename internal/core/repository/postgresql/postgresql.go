package core_postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
  QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
}


func NewPool(
	ctx context.Context, 
	connString string,
	) (*pgxpool.Pool, error) {
  return pgxpool.New(ctx, connString)
}