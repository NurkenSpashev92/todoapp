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
}

func NewUserService(userRepository UserRepositoryInterface) *UserSerivce {
	return &UserSerivce{userRepository: userRepository}
}
