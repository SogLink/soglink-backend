package refreshtoken

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type RefreshTokenRepo interface {
	Get(ctx context.Context, refreshToken string) (*entity.RefreshToken, error)
	Create(ctx context.Context, m *entity.RefreshToken) error
	Delete(ctx context.Context, refreshToken string) error
}
