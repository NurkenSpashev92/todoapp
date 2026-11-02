package tasks_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *TaskService) GetTaskById(ctx context.Context, id int) (domain.Task, error) {
	task, err := s.taskRepository.GetTaskById(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("faild to get task from repository: %w", err)
	}

	return task, nil
}
