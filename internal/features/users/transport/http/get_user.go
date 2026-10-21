package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/nurkenspashev92/todoapp/internal/core/transport/http/utls"
)

type GetUserResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed get 'limit'/'offset' query params",
		)
		return
	}

	userDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users in service",
		)
		return
	}

	response := GetUserResponse(usersDTOFromDomain(userDomains))

	responseHandler.JSONRender(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
