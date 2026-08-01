package config

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(cfg DatabaseConfig) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(
		context.Background(),
		cfg.URL,
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}