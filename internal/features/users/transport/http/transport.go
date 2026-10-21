package users_transport_http

import (
	"context"
	"net/http"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	userService UserSerivce
}

type UserSerivce interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
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
