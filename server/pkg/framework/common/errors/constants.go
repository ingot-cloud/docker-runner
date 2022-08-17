package errors

import (
	"docker-runner/pkg/framework/common/status"
	"fmt"
	"net/http"
)

var (
	ErrUnknown = New(http.StatusInternalServerError, status.InternalServerError, status.StatusText(status.InternalServerError))
)

// BadRequest error
func BadRequest(message string) error {
	return New(http.StatusBadRequest, status.BadRequest, message)
}

// Unauthorized error
func Unauthorized(message string) error {
	return New(http.StatusUnauthorized, status.Unauthorized, message)
}

// Forbidden error
func Forbidden(message string) error {
	return New(http.StatusForbidden, status.Forbidden, message)
}

// NotFound for http resource not found 404
func NotFound(path string) error {
	return New(http.StatusNotFound, status.NotFound, fmt.Sprintf("Path [%s] not found", path))
}

// MethodNotAllowed for http method not allow 405
func MethodNotAllowed(method string) error {
	return New(http.StatusMethodNotAllowed, status.MethodNotAllowed, fmt.Sprintf("Method [%s] not allow", method))
}

// IllegalOperation 非法操作
func IllegalOperation(message string) error {
	return New(http.StatusBadRequest, status.IllegalOperation, message)
}

// IllegalParameter 非法参数
func IllegalParameter(message string) error {
	return New(http.StatusBadRequest, status.IllegalParameter, message)
}
