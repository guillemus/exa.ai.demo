package db

import (
	"context"
	"fmt"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type Pool struct {
	pool *sqlitex.Pool
}

func Connect(dbPath string) (*Pool, error) {
	pool, err := sqlitex.NewPool(dbPath, sqlitex.PoolOptions{
		PoolSize:    10,
		PrepareConn: prepareConn,
	})
	if err != nil {
		return nil, fmt.Errorf("open db pool: %w", err)
	}

	return &Pool{pool: pool}, nil
}

func (x *Pool) WithConn(ctx context.Context, fn func(*sqlite.Conn) error) error {
	conn, err := x.pool.Take(ctx)
	if err != nil {
		return fmt.Errorf("take conn: %w", err)
	}
	defer x.pool.Put(conn)

	return fn(conn)
}

func (x *Pool) Close() error {
	return x.pool.Close()
}

func prepareConn(conn *sqlite.Conn) error {
	for _, stmt := range []string{
		"PRAGMA journal_mode = WAL;",
		"PRAGMA busy_timeout = 5000;",
		"PRAGMA synchronous = NORMAL;",
		"PRAGMA cache_size = -262144;",
		"PRAGMA foreign_keys = ON;",
		"PRAGMA temp_store = MEMORY;",
	} {
		if err := sqlitex.Execute(conn, stmt, nil); err != nil {
			return fmt.Errorf("prepare conn: %w", err)
		}
	}

	return nil
}
