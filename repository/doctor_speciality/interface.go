package doctorspeciality

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Doctor_specialty, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Doctor_specialty, error)
	Create(ctx context.Context, u *entity.Doctor_specialty) error
	Update(ctx context.Context, u *entity.Doctor_specialty) error
	Delete(ctx context.Context, params map[string]string) error
}
