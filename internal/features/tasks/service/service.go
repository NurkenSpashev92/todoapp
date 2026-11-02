package tasks_service

import (
	tasks_repository "github.com/nurkenspashev92/todoapp/internal/features/tasks/repository"
)

type TaskService struct {
	taskRepository tasks_repository.TaskRepositoryInterface
}

func NewTaskService(taskRepository tasks_repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}
