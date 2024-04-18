package config

import "os"

type PostgresConfig struct {
	User     string
	Password string
	Port     string
	Database string
	Host     string
}

type Config struct {
	Postgres PostgresConfig
}

func NewConfig() *Config {
	return &Config{
		Postgres: PostgresConfig{
			Database: GetEnv("POSTGRES_DATABASE", "clodoaldo"),
			Host:     GetEnv("POSTGRES_HOST", "localhost"),
			Port:     GetEnv("POSTGRES_PORT", "5432"),
			User:     GetEnv("POSTGRES_USER", "user"),
			Password: GetEnv("POSTGRES_PASSWORD", "user123"),
		},
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
