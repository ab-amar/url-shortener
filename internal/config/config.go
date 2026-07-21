package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port string
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
	
	return Config{
		Port: port,
	}, nil
}

