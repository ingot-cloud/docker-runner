package response

// D is a shortcut for map[string]interface{}
type D map[string]interface{}

// R is the Response struct
type R struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// P is a shortcut for Pagination
type P struct {
	Records interface{} `json:"records"`
	Total   int         `json:"total"`
	Current int         `json:"current"`
}
