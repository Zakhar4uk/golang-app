package tasks_transport

import (
	"net/http"
	"time"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_request "github.com/Zakhar4uk/golang-app/internal/core/transport/http/request"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompleteAt   *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func (h *TasksHTTPHandler) CreateTask(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponceHandler(log, rw)
	var request CreateTaskRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	taskDomain := domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create task",
		)
		return
	}
	response := taskDTOFromDomain(taskDomain)

	responseHandler.JSONResponce(response, http.StatusCreated)
}

func taskDTOFromDomain(task domain.Task) CreateTaskResponse {
	return CreateTaskResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompleteAt:   task.CompleteAt,
		AuthorUserID: task.AuthorID,
	}
}
