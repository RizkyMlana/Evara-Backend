package database

import (
	"context"
	"evara-backend/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(
		context.Background(),
		cfg.URL,
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}