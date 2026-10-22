package users_transport_http

import (
	"net/http"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	log.Debug("invoke CreateUser handler")

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	user := dominFromDTO(request)
	userDomain, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create in service")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func dominFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
