package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    slog.Level
}

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func ParseEnv() (*Config, error) {
	// Ignore error because in production there will be no .env file, env vars will be passed
	// in at runtime via docker run command/docker-compose
	_ = godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return nil, errors.New("SERVER_PORT environment variable is not set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL environment variable is not set")
	}

	logLevelString := os.Getenv("LOG_LEVEL")
	if logLevelString == "" {
		return nil, errors.New("LOG_LEVEL environment variable is not set")
	}
	logLevel, ok := logLevelMap[logLevelString]
	if !ok {
		return nil, errors.New("LOG_LEVEL should be one of debug|info|warning|error")
	}

	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
		LogLevel:    logLevel,
	}, nil
}
