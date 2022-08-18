package server

import (
	"context"

	"github.com/alibaba/ioc-golang/extension/config"
)

// Order Http server 顺序
const Order = 10

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:allimpls:interface=docker-runner/pkg/framework/ioc.ModuleConfigurer
// HTTPServer http web server
type HTTPServer struct {
	Context context.Context

	// Mode :debug, test, release
	Mode *config.ConfigString `config:",ingot.server.mode"`
	// Address optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Address *config.ConfigString `config:",ingot.server.address"`
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout *config.ConfigInt64 `config:",ingot.server.readTimeout"`
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout *config.ConfigInt64 `config:",ingot.server.writeTimeout"`
	// Prefix is server path prefix
	Prefix string `config:",ingot.server.prefix"`
}

// Order 执行顺序，从小到大
func (s *HTTPServer) Order() uint {
	return Order
}

// Configure 配置
func (s *HTTPServer) Configure() {}

// Run 执行
func (s *HTTPServer) Run() (func(), error) {
	return nil, nil
}
