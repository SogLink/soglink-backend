package specialty

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Specialty, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Specialty, error)
}
