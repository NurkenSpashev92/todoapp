package users_service

import (
	"context"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

var _ UserServiceInterface = (*UserSerivce)(nil)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	DeleteUserById(ctx context.Context, id int) error
	PatchUserById(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
}
