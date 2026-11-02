package users_transport_http

import (
	"net/http"

	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
	users_service "github.com/nurkenspashev92/todoapp/internal/features/users/service"
)

type UsersHTTPHandler struct {
	userService users_service.UserServiceInterface
}

func NewUsersHTTPHander(userService users_service.UserServiceInterface) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: userService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUserById,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUserById,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUserById,
		},
	}
}
