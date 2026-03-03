package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://summit:summit_dev@localhost:5432/summit?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", ""),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
