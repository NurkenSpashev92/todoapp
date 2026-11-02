package tasks_service

import (
	"context"
	"fmt"
)

func (s *TaskService) DeleteTaskById(ctx context.Context, id int) error {
	if err := s.taskRepository.DeleteTaskById(ctx, id); err != nil {
		return fmt.Errorf("faild to delete task from repository: %w", err)
	}

	return nil
}
