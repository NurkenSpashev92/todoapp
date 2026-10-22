package users_transport_http

import (
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

type UpdateUserRequest struct {
	FullName    string  `json:"full_name" validate:"omitempty,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type UpdatedUserResponse UserDTOResponse

func (h *UsersHTTPHandler) UpdateUserById(rw http.ResponseWriter, r *http.Request) {
	var request UpdateUserRequest

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	userId, err := getUserIdQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"faild to get user_id in service",
		)
		return
	}

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	user := dominFromDTO(CreateUserRequest(request))

	userDomain, err := h.userService.UpdateUserById(ctx, userId, user)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user in service",
		)
		return
	}

	response := UpdatedUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
