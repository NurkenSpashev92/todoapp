package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_postgres_pool "github.com/nurkenspashev92/todoapp/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/nurkenspashev92/todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
	users_postgres_repository "github.com/nurkenspashev92/todoapp/internal/features/users/repository/postgres"
	users_service "github.com/nurkenspashev92/todoapp/internal/features/users/service"
	users_transport_http "github.com/nurkenspashev92/todoapp/internal/features/users/transport/http"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("initializing postgres connection pool!")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres pool:", zap.Error(err))
		os.Exit(1)
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	userRepository := users_postgres_repository.NewUserRepository(pool)
	userService := users_service.NewUserService(userRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHander(userService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegiserRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegiserAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
