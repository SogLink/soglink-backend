package usecase

import (
	"context"
	// "errors"
	"time"

	"github.com/SogLink/soglink-backend/entity"
	errorspkg "github.com/SogLink/soglink-backend/errors"
	"github.com/SogLink/soglink-backend/pkg/validation"
	"github.com/SogLink/soglink-backend/repository/user"
)

type User interface {
	UserExists(ctx context.Context, guid string) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, *entity.User, error)
	GetUser(ctx context.Context, filter map[string]string) (*entity.User, error)
	CreateUser(ctx context.Context, req *entity.User) (string, error)
	ListUsers(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.User, error)
	UpdateUser(ctx context.Context, req *entity.User) error
	DeleteUser(ctx context.Context, guid string) error
}

type userUsecase struct {
	BaseUsecase
	usersRepo  user.Repository
	ctxTimeout time.Duration
}

func NewUserUsecase(ctxTimeout time.Duration, usersRepo user.Repository) User {
	return &userUsecase{
		usersRepo:  usersRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (u userUsecase) UserExists(ctx context.Context, guid string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var isExists bool = true
	_, err := u.usersRepo.Get(ctx, map[string]string{"guid": guid})
	if err != nil {
		if err == errorspkg.ErrorNotFound {
			isExists = false
		} else {
			return false, err
		}
	}

	return isExists, nil

}
func (u userUsecase) EmailExists(ctx context.Context, username string) (bool, *entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var isExists bool = true
	user, err := u.usersRepo.Get(ctx, map[string]string{"email": username})
	if err != nil {
		if err == errorspkg.ErrorNotFound {
			isExists = false
		} else {
			return false, nil, err
		}
	}

	return isExists, user, nil

}
func (u userUsecase) GetUser(ctx context.Context, filter map[string]string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.Get(ctx, filter)
}
func (u userUsecase) CreateUser(ctx context.Context, req *entity.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	passwordHash, err := validation.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	req.Password = passwordHash

	u.BaseUsecase.beforeCreate(&req.GUID, &req.CreatedAt, &req.UpdatedAt)

	return req.GUID, u.usersRepo.Create(ctx, req)
}
func (u userUsecase) ListUsers(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.List(ctx, limit, offset, filter)
}
func (u userUsecase) UpdateUser(ctx context.Context, req *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	passwordHash, err := validation.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = passwordHash

	u.BaseUsecase.beforeCreate(nil, nil, &req.UpdatedAt)

	return u.usersRepo.Update(ctx, req)
}
func (u userUsecase) DeleteUser(ctx context.Context, guid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.Delete(ctx, map[string]string{"guid": guid})
}
