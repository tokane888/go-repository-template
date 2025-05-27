package main

import (
	"errors"
	"log"

	// TODO: import元調整
	"github.com/tokane888/go-repository-template/configs"
	common "github.com/tokane888/go_common_module/v2"
	"go.uber.org/zap"
)

// アプリのversion。デフォルトは開発版。cloud上ではbuild時に-ldflagsフラグ経由でバージョンを埋め込む
var version = "dev"

func main() {
	cfg, err := configs.LoadConfig(version)
	if err != nil {
		log.Println("failed to load config:", err)
		return
	}
	logger, err := common.NewLogger(cfg.Logger)
	if err != nil {
		// zap loggerの初期化に失敗した場合のエラーハンドリング
		// zapを使用できないため、標準のlogパッケージを使用
		log.Println("failed to initialize logger:", err)
		return
	}
	defer logger.Sync()

	logger.Info("sample info")
	logger.Info("additional field sample", zap.String("key", "value"))
	logger.Warn("sample warn")
	logger.Error("sample error")
	err = errors.New("errorのサンプル")
	logger.Error("DB Connection failed", zap.Error(err))
}
