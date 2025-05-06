package main

import (
	"errors"

	"github.com/tokane888/go-repository-template/configs"
	common "github.com/tokane888/go_common_module/v2"
	"go.uber.org/zap"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return
	}
	logger, err := common.NewLogger(cfg.Logger)
	if err != nil {
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
