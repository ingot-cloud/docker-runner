package errors

// New 返回IngotError
func New(httpStatus int, statusCode, message string) error {
	return &IngotError{
		HttpStatusCode: httpStatus,
		StatusCode:     statusCode,
		Message:        message,
	}
}

// IngotError 自定义异常，包含Http状态码，自定义状态码和消息
type IngotError struct {
	HttpStatusCode int
	StatusCode     string
	Message        string
}

// Error error interface
func (e *IngotError) Error() string {
	return e.Message
}
