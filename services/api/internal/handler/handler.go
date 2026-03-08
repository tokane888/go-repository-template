package handler

import (
	"log/slog"

	"github.com/tokane888/go-repository-template/services/api/internal/usecase"
)

type Handler struct {
	logger      *slog.Logger
	userUseCase usecase.UserUseCase
}

func NewHandler(logger *slog.Logger, userUseCase usecase.UserUseCase) *Handler {
	return &Handler{
		logger:      logger,
		userUseCase: userUseCase,
	}
}
