package user_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	core_http_request "github.com/Zakhar4uk/golang-app/internal/core/transport/http/request"
	core_http_response "github.com/Zakhar4uk/golang-app/internal/core/transport/http/response"
	core_http_types "github.com/Zakhar4uk/golang-app/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" swaggertype:"string"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("'FullName' can't be NULL")
		}
		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("'FullName' must be between 3 and 100 symbols")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("'PhoneNumber' must be between 10 and 15 symbos")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("'PhoneNumber' must starts witch '+' symbol")
			}
		}
	}
	return nil
}

type PatchUserResponse UserDTOResponse

// @PatchUser godoc
// @Summary Измненение пользователя
// @Description Измненение информации о уже существующем пользователе
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID изменяемого пользователя"
// @Param request body PatchUserRequest true "PatchUser тело запроса"
// @Success 200 {object} PatchUserRequest "Успешное измненение"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/id [patch]
func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponceHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user ID path value")
		return
	}

	var request PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and vaidate HTTP request",
		)
		return
	}

	userPatch := userPatchFromRequest(request)
	userDomain, err := h.UserService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponce(response, http.StatusOK)

}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
