package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}

	DB = pool
}