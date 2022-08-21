package web

import "github.com/gin-gonic/gin"

// APIFunc 接口处理器方法, 返回响应结果
type APIFunc func(*Context) (interface{}, error)

// MiddlewareFunc 中间件处理器方法
type MiddlewareFunc func(*Context) error

// apiToHandlerFun ApiFunc to gin.HandlerFunc
func apiToHandlerFun(api APIFunc) gin.HandlerFunc {
	return func(ginCon *gin.Context) {
		con := New(ginCon)
		result, err := api(con)
		if con.isResult {
			return
		}
		if err != nil {
			con.Error(err)
			return
		}
		if result != nil {
			con.OK(result)
			return
		}
		con.Empty()
	}
}

// middlewareToHandlerFunc MiddlewareFunc to gin.HandlerFunc
func middlewareToHandlerFunc(middleware MiddlewareFunc) gin.HandlerFunc {
	return func(ginCon *gin.Context) {
		con := New(ginCon)
		err := middleware(con)

		if con.isResult {
			if !con.IsAborted() {
				con.Abort()
			}
			return
		}

		if err != nil {
			con.Error(err)
			if !con.IsAborted() {
				con.Abort()
			}
			return
		}
	}
}
