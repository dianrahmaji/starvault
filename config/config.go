package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IDNApiUrl string
	DSN       string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	IDNApiUrl := os.Getenv("IDN_API_URL")

	postgresDB := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", postgresUser, postgresPassword, postgresDB)

	return &Config{
		IDNApiUrl: IDNApiUrl,
		DSN:       dsn,
	}, nil
}
