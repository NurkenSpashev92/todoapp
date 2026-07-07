package users_transport_http

import (
	"net/http"

	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	userService UserSerivce
}

type UserSerivce interface {
}

func NewUsersHTTPHander(userService UserSerivce) *UsersHTTPHandler {
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
	}
}
