package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/nurkenspashev92/todoapp/internal/core/transport/http/utls"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUserById(rw http.ResponseWriter, r *http.Request) {
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

	userDomain, err := h.userService.GetUserById(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user in service",
		)
		return
	}

	response := GetUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdQueryParam(r *http.Request) (int, error) {
	userId, err := core_http_utils.GetIntPathValue(r, "user_id")
	if err != nil {
		return 0, fmt.Errorf("get `user_id` query param: %w", err)
	}

	return userId, nil
}
