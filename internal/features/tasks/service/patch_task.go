package tasks_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *TaskService) PatchTaskById(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error) {
	task, err := s.taskRepository.GetTaskById(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("task not found: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	patchedTask, err := s.taskRepository.PatchTaskById(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("faild to update task in repository: %w", err)
	}

	return patchedTask, nil
}
