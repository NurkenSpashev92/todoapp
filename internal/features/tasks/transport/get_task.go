package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTaskById(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	taskId, err := getTaskIdPathValue(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task id from path",
		)
		return
	}

	taskDomain, err := h.taskService.GetTaskById(ctx, taskId)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task in service",
		)
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getTaskIdPathValue(r *http.Request) (int, error) {
	taskId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		return 0, fmt.Errorf("get `id` path value: %w", err)
	}

	return taskId, nil
}
