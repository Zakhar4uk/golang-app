package user_transport_http

import (
	"net/http"

	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_request "github.com/Zakhar4uk/golang-app/internal/core/transport/http/request"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (s *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponceHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}
	user, err := s.UserService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}
	respose := GetUserResponse(userDTOFromDomain(user))
	responseHandler.JSONResponce(respose, http.StatusOK)
}
