package handler

import (
	"context"

	"github.com/SogLink/soglink-backend/api/middleware"
	"github.com/SogLink/soglink-backend/pkg/config"
	usecase "github.com/SogLink/soglink-backend/services"
	"go.uber.org/zap"
)

type HandlerArguments struct {
	ReshreshTokenUsecase usecase.RefreshToken
	PatientUsecase       usecase.Patient
	UserUsecase          usecase.User
	Logger               *zap.Logger
	Config               *config.Config
}

type BaseHandler struct{}

func (h *BaseHandler) GetAuthData(ctx context.Context) (map[string]string, bool) {
	data, ok := ctx.Value(middleware.CtxKeyAuthData).(map[string]string)
	return data, ok
}
