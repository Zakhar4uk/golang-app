package user_transport_http

import (
	"net/http"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_request "github.com/Zakhar4uk/golang-app/internal/core/transport/http/request"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validae:"omitemty,min=10,max=15,startswith=+"`
}
type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponceHandler(log, rw)
	var request CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responceHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.USerService.CreateUser(ctx, userDomain)
	if err != nil {
		responceHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responce := CreateUserResponse(userDTOFromDomain(userDomain))
	responceHandler.JSONResponce(responce, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
