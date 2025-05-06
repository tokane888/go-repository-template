package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	common "github.com/tokane888/go_common_module/v2"
)

var version = "dev" // アプリのversion。デフォルトは開発版。cloud上ではbuild時に上書き

type Config struct {
	Env    string
	Logger common.LoggerConfig
	// 必要に応じてDatabaseConfig等各structへ注入する設定追加
}

// LoadConfig loads environment variables into Config
func LoadConfig() (*Config, error) {
	env := getEnv("ENV", "local")
	envFile := fmt.Sprintf(".env/.env.%s", env)
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", envFile, err)
	}

	cfg := &Config{
		Env: env,
		Logger: common.LoggerConfig{
			AppName:    getEnv("APP_NAME", ""),
			AppVersion: version,
			Level:      getEnv("LOG_LEVEL", "info"),
			Format:     getEnv("LOG_FORMAT", "local"),
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
