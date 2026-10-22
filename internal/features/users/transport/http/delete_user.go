package users_transport_http

import (
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUserById(rw http.ResponseWriter, r *http.Request) {
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

	if err := h.userService.DeleteUserById(ctx, userId); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user in service",
		)
		return
	}

	responseHandler.NoContentResponse()
}
