package tasks_transport_http

import "github.com/nurkenspashev92/todoapp/internal/core/domain"

type TaskDTOResponse struct {
	ID           int     `json:"id"`
	Version      int     `json:"version"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	Completed    bool    `json:"completed"`
	AuthorUserId int     `json:"author_user_id"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		AuthorUserId: task.AuthorUserId,
	}
}

func tasksDTOFromDomain(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}

	return tasksDTO
}
