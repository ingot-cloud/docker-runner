package web

import (
	ingotErrors "docker-runner/pkg/framework/common/errors"
	"docker-runner/pkg/framework/common/model/response"
	"docker-runner/pkg/framework/common/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New 实例Context
func New(gin *gin.Context) *Context {
	return &Context{
		Context:  gin,
		isResult: false,
	}
}

// Context 包装gin.Context
type Context struct {
	*gin.Context
	isResult bool
}

// OK 响应成功
func (con *Context) OK(data interface{}) {
	con.Result(http.StatusOK, status.OK, data, status.StatusText(status.OK))
}

// Empty 响应成功
func (con *Context) Empty() {
	con.Result(http.StatusOK, status.OK, response.D{}, status.StatusText(status.OK))
}

// Error 响应失败
func (con *Context) Error(err error) {
	ingotErr := ingotErrors.Unpack(err)
	con.Result(ingotErr.HttpStatusCode, ingotErr.StatusCode, response.D{}, ingotErr.Message)
}

// Result 响应结果
func (con *Context) Result(statusCode int, code string, data interface{}, message string) {
	con.isResult = true
	con.JSON(statusCode, response.R{
		Code:    code,
		Data:    data,
		Message: message,
	})
}
