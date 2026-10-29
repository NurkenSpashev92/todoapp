package tasks_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *TaskService) CreateUser(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate taks doamin: %w", err)
	}

	task, err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create user in repository: %w", err)
	}

	return task, nil
}
