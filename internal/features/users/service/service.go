package users_service

import (
	"context"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type UserSerivce struct {
	userRepository UserRepositoryInterface
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	DeleteUserById(ctx context.Context, id int) error
	UpdateUserById(ctx context.Context, id int, user domain.User) (domain.User, error)
}

func NewUserService(userRepository UserRepositoryInterface) *UserSerivce {
	return &UserSerivce{userRepository: userRepository}
}
