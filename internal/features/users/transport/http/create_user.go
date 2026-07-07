package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}
}

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write([]byte("Hello world")); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
