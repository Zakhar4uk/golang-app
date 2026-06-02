package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	h.rw.WriteHeader(statusCode)

	responce := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(responce); err != nil {
		h.log.Error("write HTTP responce", zap.Error(err))
	}
}
