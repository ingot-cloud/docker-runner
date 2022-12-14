//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package server

import (
	ioc "docker-runner/pkg/framework/ioc"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &hTTPServer_{}
		},
	})
	hTTPServerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &HTTPServer{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"allimpls": map[string]interface{}{
					"interfaces": []interface{}{
						new(ioc.ModuleConfigurer),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(hTTPServerStructDescriptor)
}

type hTTPServer_ struct {
	Order_ func() uint
	Run_   func() (func(), error)
}

func (h *hTTPServer_) Order() uint {
	return h.Order_()
}

func (h *hTTPServer_) Run() (func(), error) {
	return h.Run_()
}

type HTTPServerIOCInterface interface {
	Order() uint
	Run() (func(), error)
}

var _hTTPServerSDID string
