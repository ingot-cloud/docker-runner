package status

// Standard status
const (
	Success             = "S200"
	BadRequest          = "S400"
	Unauthorized        = "S401"
	Forbidden           = "S403"
	NotFound            = "S404"
	MethodNotAllowed    = "S405"
	InternalServerError = "S500"

	IllegalOperation = "S001"
	IllegalParameter = "S002"
)

var statusText = map[string]string{
	Success:             "OK",
	BadRequest:          "Bad Request",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not Found",
	MethodNotAllowed:    "Method Not Allowed",
	InternalServerError: "Internal Server Error",

	IllegalOperation: "Illegal Operation",
	IllegalParameter: "Illegal Parameter",
}

// StatusText 返回状态码文本
func StatusText(code string) string {
	return statusText[code]
}
