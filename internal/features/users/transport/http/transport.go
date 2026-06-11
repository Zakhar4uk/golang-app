package user_transport_http

import (
	"context"
	"net/http"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_http_server "github.com/Zakhar4uk/golang-app/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	USerService USerService
}

type USerService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
}

func NewUsersHTTPHanlder(USerService USerService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		USerService: USerService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
