package file

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.File, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.File, error)
	Create(ctx context.Context, u *entity.File) error
	Update(ctx context.Context, u *entity.File) error
	Delete(ctx context.Context, params map[string]string) error
}
