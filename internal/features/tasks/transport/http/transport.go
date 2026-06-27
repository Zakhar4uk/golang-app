package tasks_transport

import (
	"context"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_http_server "github.com/Zakhar4uk/golang-app/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

func NewTasksHTTPHandler(
	tasksService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{}
}
