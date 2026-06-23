package user_transport_http

import (
	"net/http"

	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
	core_http_utils "github.com/Zakhar4uk/golang-app/internal/core/transport/http/utils"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponceHandler(log, rw)
	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}
	if err := h.UserService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user",
		)
		return
	}
	responseHandler.NoContentResponse()
}
