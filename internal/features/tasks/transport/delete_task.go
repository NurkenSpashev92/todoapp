package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) DeleteTaskById(rw http.ResponseWriter, r *http.Request) {
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

	if err := h.taskService.DeleteTaskById(ctx, taskId); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete task in service",
		)
		return
	}

	responseHandler.NoContentResponse()
}
