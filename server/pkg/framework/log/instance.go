package log

import (
	ioc "docker-runner/pkg/framework/ioc"

	"github.com/alibaba/ioc-golang/extension/config"
	"github.com/sirupsen/logrus"
)

var (
	std = logrus.New()
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:allimpls:interface=docker-runner/pkg/framework/ioc.ModuleConfigurer
// LogConfigurer 日志配置
type LogConfigurer struct {
	Level         *config.ConfigInt    `config:",ingot.log.level"`
	Format        *config.ConfigString `config:",ingot.log.format"`
	Output        *config.ConfigString `config:",ingot.log.output"`
	OutputFileDir *config.ConfigString `config:",ingot.log.outputFileDir"`
	LogSoftLink   *config.ConfigString `config:",ingot.log.logSoftLink"`
}

// Order 排序
func (c *LogConfigurer) Order() uint {
	return 0
}

// Configure 配置方法
func (c *LogConfigurer) Configure() {
	c.setLevel(c.Level.Value())
	c.setFormatter(c.Format.Value())
}

// Run 执行
func (c *LogConfigurer) Run() (func(), error) {
	return ioc.EmptyCleanFn, nil
}

// SetLevel 设定日志级别
func (c *LogConfigurer) setLevel(level int) {
	std.SetLevel(logrus.Level(level))
}

// SetFormatter 设定日志输出格式
func (c *LogConfigurer) setFormatter(format string) {
	switch format {
	case "json":
		std.SetFormatter(new(logrus.JSONFormatter))
	default:
		// ReportCaller 开启后性能降低
		// logrus.SetReportCaller(true)
		std.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}
