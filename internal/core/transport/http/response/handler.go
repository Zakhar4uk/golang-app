package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/Zakhar4uk/golang-app/internal/core/errors"
	core_logger "github.com/Zakhar4uk/golang-app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponceHandler(
	log *core_logger.Logger,
	rw http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) PanicResponce(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("enexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponce(statusCode, err, msg)
}

func (h *HTTPResponseHandler) JSONResponce(
	responceBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)
	if err := json.NewEncoder(h.rw).Encode(responceBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponce(statusCode, err, msg)

}

func (h *HTTPResponseHandler) errorResponce(
	statusCode int,
	err error,
	msg string,
) {

	responce := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponce(responce, statusCode)
}
