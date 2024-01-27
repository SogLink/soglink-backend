package doctor

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Doctor, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Doctor, error)
	Create(ctx context.Context, u *entity.Doctor) error
	Update(ctx context.Context, u *entity.Doctor) error
	Delete(ctx context.Context, params map[string]string) error
}
