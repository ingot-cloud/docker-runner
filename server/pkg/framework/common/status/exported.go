package status

const (
	Success             = "0200"
	BadRequest          = "0400"
	Unauthorized        = "0401"
	Forbidden           = "0403"
	NotFound            = "0404"
	MethodNotAllowed    = "0405"
	InternalServerError = "0500"
)

var statusText = map[string]string{
	Success:             "Success",
	BadRequest:          "",
	Unauthorized:        "",
	Forbidden:           "",
	NotFound:            "",
	MethodNotAllowed:    "",
	InternalServerError: "未知错误",
}

// StatusText 返回状态码文本
func StatusText(code string) string {
	return statusText[code]
}
