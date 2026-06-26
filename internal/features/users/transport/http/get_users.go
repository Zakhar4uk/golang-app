package user_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_request "github.com/Zakhar4uk/golang-app/internal/core/transport/http/request"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter,
	r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponceHandler(log, rw)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query param",
		)
		return
	}
	userDomains, err := h.UserService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))

	responseHandler.JSONResponce(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {

	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get 'limit' query param: %w", err,
		)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get 'offset' query param: %w", err,
		)
	}
	return limit, offset, nil
}
