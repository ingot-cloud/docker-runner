package log

import (
	"docker-runner/pkg/framework/common/utils"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/alibaba/ioc-golang/extension/config"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
	var file *os.File
	if c.Output.Value() != "" {
		switch c.Output.Value() {
		case "stdout":
			SetOutput(os.Stdout)
		case "stderr":
			SetOutput(os.Stderr)
		case "file":
			if dir := c.OutputFileDir.Value(); dir != "" {
				fileWriter, err := c.rotate()
				if err != nil {
					return func() {}, err
				}
				gin.DefaultWriter = io.MultiWriter(fileWriter, os.Stdout)
				SetOutput(fileWriter)
			}
		}
	}
	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
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

func (c *LogConfigurer) rotate() (io.Writer, error) {
	if ok, _ := utils.PathExists(c.OutputFileDir.Value()); !ok {
		// directory not exist
		fmt.Println("create log directory")
		err := os.Mkdir(c.OutputFileDir.Value(), os.ModePerm)
		if err != nil {
			fmt.Println("mkdir error - ", err)
		}
	}
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s%s", c.OutputFileDir.Value(), string(os.PathSeparator), "%Y-%m-%d-%H-%M.log"),
		// generate soft link, point to latest log file
		rotatelogs.WithLinkName(c.LogSoftLink.Value()),
		// maximum time to save log files
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// time period of log file switching
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	return writer, nil
}
