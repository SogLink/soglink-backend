package location

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Location, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Location, error)
	Create(ctx context.Context, u *entity.Location) error
	Update(ctx context.Context, u *entity.Location) error
	Delete(ctx context.Context, params map[string]string) error
}
