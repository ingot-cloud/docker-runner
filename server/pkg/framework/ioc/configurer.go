package ioc

func EmptyCleanFn() {}

// ModuleConfigurer 模块配置器
type ModuleConfigurer interface {
	// Order 执行顺序，从小到大
	Order() uint

	// Configure 配置
	Configure()

	// Run 执行
	Run() (func(), error)
}
