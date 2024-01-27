package patient

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Patient, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Patient, error)
	Create(ctx context.Context, u *entity.Patient) error
	Update(ctx context.Context, u *entity.Patient) error
	Delete(ctx context.Context, params map[string]string) error
}
