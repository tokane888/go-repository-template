package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tokane888/go-repository-template/pkg/logger"
)

// Config reads environment variables and holds per-struct configuration.
type Config struct {
	Env    string
	Logger logger.Config
	// Add configs for injection into structs such as DatabaseConfig as needed.
}

// NewConfig loads environment variables into Config
func NewConfig(version string) (*Config, error) {
	env := getEnv("ENV", "local")
	envFile := ".env/.env." + env
	if env == "local" {
		err := godotenv.Load(envFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", envFile, err)
		}
	}

	cfg := &Config{
		Env: env,
		Logger: logger.Config{
			AppName:    getEnv("APP_NAME", ""),
			AppVersion: version,
			Env:        env,
			Level:      getEnv("LOG_LEVEL", "info"),
			Format:     getEnv("LOG_FORMAT", "cloud"),
		},
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
