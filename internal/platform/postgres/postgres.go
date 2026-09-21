// Package postgres provides a connection pool to the PostgreSQL database
// shared by every repository in the application.
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool opens a connection pool using the given DSN
// (e.g. "postgres://user:pass@host:5432/db?sslmode=disable").
func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to create pool: %w", err)
	}

	// Ping right away so startup fails fast if the database is unreachable,
	// instead of only surfacing the problem on the first real query.
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres: failed to ping database: %w", err)
	}

	return pool, nil
}
