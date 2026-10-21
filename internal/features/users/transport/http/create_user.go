package users_transport_http

import (
	"net/http"

	"go.uber.org/zap"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest

	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		log.Error("error decode json body", zap.String("message", err.Error()))
		http.Error(rw, "invalid json body", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
}

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")

	if _, err := rw.Write([]byte("Hello world")); err != nil {
		log.Error("error decode json body", zap.String("message", err.Error()))
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
}
