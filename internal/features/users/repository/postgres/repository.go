package users_postgres_repository

import (
	core_postgres_pool "github.com/nurkenspashev92/todoapp/internal/core/repository/postgres/pool"
	users_repository "github.com/nurkenspashev92/todoapp/internal/features/users/repository"
)

var _ users_repository.UserRepositoryInterface = (*UserRepository)(nil)

type UserRepository struct {
	pool core_postgres_pool.Pool
}

func NewUserRepository(pool core_postgres_pool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
