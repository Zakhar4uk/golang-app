package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Zakhar4uk/golang-app/docs"
	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_pgx_pool "github.com/Zakhar4uk/golang-app/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Zakhar4uk/golang-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Zakhar4uk/golang-app/internal/core/transport/http/server"
	statistics_posgres_repository "github.com/Zakhar4uk/golang-app/internal/features/statistics/repository/posgres"
	statistics_service "github.com/Zakhar4uk/golang-app/internal/features/statistics/service"
	statistics_tranposrt_http "github.com/Zakhar4uk/golang-app/internal/features/statistics/transport/http"
	tasks_posgres_repository "github.com/Zakhar4uk/golang-app/internal/features/tasks/repository/posgres"
	tasks_service "github.com/Zakhar4uk/golang-app/internal/features/tasks/service"
	tasks_transport_http "github.com/Zakhar4uk/golang-app/internal/features/tasks/transport/http"
	users_posgres_repository "github.com/Zakhar4uk/golang-app/internal/features/users/repository/posgres"
	users_service "github.com/Zakhar4uk/golang-app/internal/features/users/service"
	user_transport_http "github.com/Zakhar4uk/golang-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

// @title        Golang Todo API
// @version      1.0
// @description  Todo Applicaton REST-API scheme
// @host         127.0.0.1:5050
// @BasePath     /api/v1
func main() {
	time.Local = timeZone

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

	logger.Debug("application time zone", zap.Any("zone", timeZone))

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

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_posgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_posgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_tranposrt_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("init HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
