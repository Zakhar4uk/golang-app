package user_transport_http

import (
	"context"
	"net/http"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_http_server "github.com/Zakhar4uk/golang-app/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	UserService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
}

func NewUsersHTTPHanlder(UserService UserService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UserService: UserService,
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
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
	}
}
