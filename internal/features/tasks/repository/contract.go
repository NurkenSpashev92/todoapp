package tasks_repository

import (
	"context"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error)
	GetTaskById(ctx context.Context, id int) (domain.Task, error)
	DeleteTaskById(ctx context.Context, id int) error
	PatchTaskById(ctx context.Context, id int, task domain.Task) (domain.Task, error)
}
