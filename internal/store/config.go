package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// creates a new connection pool using pgxpool
func NewPgxpoolConn(addr string, maxConns int, maxConnIdleTime string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(maxConns)

	idleDuration, err := time.ParseDuration(maxConnIdleTime)
	if err != nil {
		return nil, err
	}
	config.MaxConnIdleTime = idleDuration

	// Create a connection pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
