package tasks_transport_http

import (
	"net/http"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string     `json:"title" validate:"required,min=1,max=1000"`
	Description  *string    `json:"description" validate:"omitempty,min=1,max=1000"`
	Completed    bool       `json:"completed" default:"false"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_At"`
	AuthorUserId int        `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	var request CreateTaskRequest

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	task := domainFromDTO(request)
	taskDomain, err := h.taskService.CreateTask(ctx, task)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task in service")
		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		dto.Title,
		dto.Description,
		dto.AuthorUserId,
	)
}
