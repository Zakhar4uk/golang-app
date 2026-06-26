package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_pgx_pool "github.com/Zakhar4uk/golang-app/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Zakhar4uk/golang-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Zakhar4uk/golang-app/internal/core/transport/http/server"
	users_posgres_repository "github.com/Zakhar4uk/golang-app/internal/features/users/repository/posgres"
	users_service "github.com/Zakhar4uk/golang-app/internal/features/users/service"
	user_transport_http "github.com/Zakhar4uk/golang-app/internal/features/users/transport/http"
	"go.uber.org/zap"
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

	logger.Debug("init postgres connection pool")

	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("init feature", zap.String("feature", "users"))
	userRepository := users_posgres_repository.NewUserRepository(pool)
	usersService := users_service.NewUserService(userRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHanlder(usersService)

	logger.Debug("init HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
