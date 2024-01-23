package appointment

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Appointment, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Appointment, error)
	Create(ctx context.Context, u *entity.Appointment) error
	Update(ctx context.Context, u *entity.Appointment) error
	Delete(ctx context.Context, params map[string]string) error
}
