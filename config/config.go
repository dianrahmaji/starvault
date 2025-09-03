package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IDNApiUrl string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	IDNApiUrl := os.Getenv("IDN_API_URL")
	if IDNApiUrl == "" {
		return nil, errors.New("missing IDN_API_URL environment")
	}

	return &Config{
		IDNApiUrl: IDNApiUrl,
	}, nil
}
