package tasks_service

import (
	"context"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

var _ TaskServiceInterface = (*TaskService)(nil)

type TaskServiceInterface interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error)
	GetTaskById(ctx context.Context, id int) (domain.Task, error)
	DeleteTaskById(ctx context.Context, id int) error
	PatchTaskById(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error)
}
