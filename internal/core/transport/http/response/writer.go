package core_http_response

import "net/http"

var (
	StatusCodeUnitialLized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponceWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUnitialLized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) GetStatusCode() int {
	if rw.statusCode == StatusCodeUnitialLized {
		return http.StatusOK
	}
	return rw.statusCode
}
