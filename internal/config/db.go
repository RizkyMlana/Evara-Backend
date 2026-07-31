package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase() (*pgxpool.Pool, error) {

	dsn := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(
		context.Background(),
		dsn,
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}