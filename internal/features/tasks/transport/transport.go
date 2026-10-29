package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	taskService TaskService
}

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

func NewTaskHTTPHandler(taskService TaskService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		taskService: taskService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
	}
}
