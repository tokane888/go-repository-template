package main

import (
	"errors"
	"log"
	"log/slog"

	// TODO: import元調整
	pkglogger "github.com/tokane888/go-repository-template/pkg/logger"
	"github.com/tokane888/go-repository-template/services/sample/internal/config"
)

// アプリのversion。デフォルトは開発版。cloud上ではbuild時に-ldflagsフラグ経由でバージョンを埋め込む
var version = "dev"

func main() {
	cfg, err := config.LoadConfig(version)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}
	logger := pkglogger.NewLogger(cfg.Logger)

	logger.Info("sample batch info")
	logger.Info("additional field sample", slog.String("key", "value"))
	logger.Warn("sample warn")
	logger.Error("sample error")
	err = errors.New("errorのサンプル")
	logger.Error("sample error with detail", slog.Any("error", err))
}
