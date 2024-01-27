package emr

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Emr, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Emr, error)
	Create(ctx context.Context, u *entity.Emr) error
	Update(ctx context.Context, u *entity.Emr) error
	Delete(ctx context.Context, params map[string]string) error
}
