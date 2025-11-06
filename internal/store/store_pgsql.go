package store

import (
	//static "app/internal/store/static"
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	upper "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type PGSQLConnectors struct {
	NativeConn   *pgxpool.Pool
	Db           *sql.DB
	Dbx          *sqlx.DB
	Ormx         upper.Session
	QueryBuilder *sq.StatementBuilderType
}

type PGSQLStore struct {
	Db *sql.DB
	*PGSQLConnectors
	//*static.Queries
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	var ormx upper.Session
	ormx, _ = postgresql.New(db)
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	connectors := &PGSQLConnectors{
		Db:           db,
		Dbx:          sqlx.NewDb(db, "postgres"),
		Ormx:         ormx,
		QueryBuilder: &queryBuilder,
	}
	return &PGSQLStore{
		Db:              db,
		PGSQLConnectors: connectors,
		//Queries:         static.New(db),
	}
}

func (store *PGSQLStore) AttachNativeConn(ctx context.Context, addr string, maxConns int, maxConnIdleTime string) error {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return fmt.Errorf("parse config error: %w", err)
	}

	config.MaxConns = int32(maxConns)

	idleDuration, err := time.ParseDuration(maxConnIdleTime)
	if err != nil {
		return fmt.Errorf("invalid idle time duration: %w", err)
	}
	config.MaxConnIdleTime = idleDuration

	// Gracefully close existing pool if already set
	if store.NativeConn != nil {
		store.NativeConn.Close()
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctxWithTimeout, config)
	if err != nil {
		return fmt.Errorf("create pool error: %w", err)
	}

	if err := pool.Ping(ctxWithTimeout); err != nil {
		pool.Close()
		return fmt.Errorf("ping failed: %w", err)
	}

	store.NativeConn = pool
	return nil
}

func (store *PGSQLStore) GetNativeConn() interface{} {
	return store.NativeConn
}

// ExecTx executes a function within a database transaction
// func (store *PGSQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.Db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }

// getInsertID returns the integer value of a newly inserted id (using upper)
func getInsertID(i upper.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
