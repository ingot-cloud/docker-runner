package boot

import (
	"docker-runner/pkg/framework/ioc"
	"sort"
)

type ModuleConfigurers []ioc.ModuleConfigurer

func (mc ModuleConfigurers) Len() int {
	return len(mc)
}

func (mc ModuleConfigurers) Less(i, j int) bool {
	return mc[i].Order() < mc[j].Order()
}

func (mc ModuleConfigurers) Swap(i, j int) {
	mc[i], mc[j] = mc[j], mc[i]
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// ModuleConfigurer 模块配置器
type ModuleConfigurer struct {
	ModuleConfigurers []ioc.ModuleConfigurer `allimpls:""`
}

// Configure 配置有所模块
func (m *ModuleConfigurer) Configure() ([]func(), error) {
	configurers := ModuleConfigurers(m.ModuleConfigurers)
	sort.Sort(configurers)
	for _, configurer := range configurers {
		configurer.Configure()
	}

	var cleanFns []func()
	for _, configurer := range configurers {
		fn, err := configurer.Run()
		if err != nil {
			return cleanFns, err
		}
		cleanFns = append(cleanFns, fn)
	}

	return cleanFns, nil
}
