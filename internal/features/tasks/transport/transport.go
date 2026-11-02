package tasks_transport_http

import (
	"net/http"

	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
	tasks_service "github.com/nurkenspashev92/todoapp/internal/features/tasks/service"
)

type TasksHTTPHandler struct {
	taskService tasks_service.TaskServiceInterface
}

func NewTaskHTTPHandler(taskService tasks_service.TaskServiceInterface) *TasksHTTPHandler {
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
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTaskById,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTaskById,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTaskById,
		},
	}
}
