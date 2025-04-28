package main

import (
	"errors"

	"github.com/tokane888/go-repository-template/config"
	common "github.com/tokane888/go_common_module"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}
	logger, err := common.NewLogger(cfg.Logger)
	if err != nil {
		return
	}

	logger.Info("sample info")
	logger.Info("additional field sample", zap.String("key", "value"))
	logger.Warn("sample warn")
	logger.Error("sample error")
	err = errors.New("errorのサンプル")
	logger.Errorw("DB Connection failed", err)
}
