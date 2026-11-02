package users_service

import (
	users_repository "github.com/nurkenspashev92/todoapp/internal/features/users/repository"
)

type UserSerivce struct {
	userRepository users_repository.UserRepositoryInterface
}

func NewUserService(userRepository users_repository.UserRepositoryInterface) *UserSerivce {
	return &UserSerivce{userRepository: userRepository}
}
