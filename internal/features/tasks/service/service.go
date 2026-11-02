package tasks_service

import (
	"context"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type TaskService struct {
	taskRepository TaskRepositoryInterface
}

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, limit *int, offset *int) ([]domain.Task, error)
	GetTaskById(ctx context.Context, id int) (domain.Task, error)
	DeleteTaskById(ctx context.Context, id int) error
	PatchTaskById(ctx context.Context, id int, task domain.Task) (domain.Task, error)
}

func NewTaskService(taskRepository TaskRepositoryInterface) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}
