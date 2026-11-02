package tasks_postgres_repository

import (
	core_postgres_pool "github.com/nurkenspashev92/todoapp/internal/core/repository/postgres/pool"
	tasks_repository "github.com/nurkenspashev92/todoapp/internal/features/tasks/repository"
)

var _ tasks_repository.TaskRepositoryInterface = (*TaskRepository)(nil)

type TaskRepository struct {
	pool core_postgres_pool.Pool
}

func NewTaskRepository(pool core_postgres_pool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}
