package clinic

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Clinic, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Clinic, error)
	Create(ctx context.Context, u *entity.Clinic) error
	Update(ctx context.Context, u *entity.Clinic) error
	Delete(ctx context.Context, params map[string]string) error
}
