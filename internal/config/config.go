package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func NewConfig() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return Config{}, err
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, errors.New("database url is required")
	}

	return Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}, nil
}
