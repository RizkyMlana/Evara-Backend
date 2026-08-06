package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)



type DBTX interface {
	Exec(
		ctx context.Context,
		sql string,
		arguments ...any,
	)(pgconn.CommandTag, error)

	Query(
		ctx context.Context,
		sql string,
		args ...any,
	) (pgx.Rows, error)

	QueryRow (
		ctx context.Context,
		sql string,
		args ...any,
	) pgx.Row
}

type TransactionManager struct {
	db *pgxpool.Pool
}

func NewTransactionManager(
	db *pgxpool.Pool,
) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (m *TransactionManager) WithinTransaction(
	ctx context.Context,
	fn func(DBTX) error,
) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}