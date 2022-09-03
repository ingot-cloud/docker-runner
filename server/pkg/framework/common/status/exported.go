package status

// Standard status
const (
	OK                  = "S0200"
	BadRequest          = "S0400"
	Unauthorized        = "S0401"
	Forbidden           = "S0403"
	NotFound            = "S0404"
	MethodNotAllowed    = "S0405"
	InternalServerError = "S0500"

	IllegalOperation = "S0001"
	IllegalParameter = "S0002"
)

var statusText = map[string]string{
	OK:                  "OK",
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
