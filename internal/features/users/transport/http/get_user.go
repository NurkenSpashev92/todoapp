package users_transport_http

import (
	"net/http"

	"go.uber.org/zap"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
)

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")

	if _, err := rw.Write([]byte("Hello world")); err != nil {
		log.Error("error decode json body", zap.String("message", err.Error()))
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
}
