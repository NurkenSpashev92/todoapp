package tasks_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func (s *TaskService) GetTasks(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvlidArgument,
		)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvlidArgument,
		)
	}

	tasks, err := s.taskRepository.GetTasks(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("faild to get tasks: %w", err)
	}

	return tasks, nil
}
