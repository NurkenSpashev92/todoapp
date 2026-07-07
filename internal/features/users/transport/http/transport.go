package users_transport_http

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
