package errors

import (
	"docker-runner/pkg/framework/common/status"
	"net/http"
)

// Unpack error
func Unpack(err error) *IngotError {
	if e, ok := err.(*IngotError); ok {
		return e
	}

	message := status.StatusText(status.InternalServerError)
	if err != nil {
		message = err.Error()
	}

	// 默认拆包异常
	return &IngotError{
		HttpStatusCode: http.StatusInternalServerError,
		StatusCode:     status.InternalServerError,
		Message:        message,
	}
}
