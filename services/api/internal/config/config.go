package config

import (
	"errors"
	"os"
)

type Config struct {
	AppPort     string
	DatabaseURL string
}

func Load() (Config, error) {
	appPort := getenv("APP_PORT", "8080")
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}

	return Config{
		AppPort:     appPort,
		DatabaseURL: databaseURL,
	}, nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
