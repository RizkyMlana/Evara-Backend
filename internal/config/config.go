package config

import (
	"fmt"
	"os"
)

type Config struct{
	Server ServerConfig
	Database DatabaseConfig
	JWT JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URL string
}

type JWTConfig struct {
	URL string
}

func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT"),
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL"),
		},

		JWT: JWTConfig{
			URL: getEnv("SUPABASE_JWKS_URL"),
		},
	}
	return cfg, nil
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("%s is not set", key))
	}
	return value
}