package attachment

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
)

type Repository interface {
	Get(ctx context.Context, params map[string]string) (*entity.Attachment, error)
	List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Attachment, error)
	Create(ctx context.Context, u *entity.Attachment) error
	Update(ctx context.Context, u *entity.Attachment) error
	Delete(ctx context.Context, params map[string]string) error
}
