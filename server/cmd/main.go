package main

import (
	"docker-runner/pkg/framework/boot"
	"docker-runner/pkg/framework/log"
	"os"

	"github.com/urfave/cli/v2"
)

var VERSION = "0.1.0"
var SERVER_NAME = ""
var USAGE = ""

func main() {
	app := &cli.App{
		Name:    SERVER_NAME,
		Version: VERSION,
		Commands: []*cli.Command{
			serverCmd(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Errorf(err.Error())
	}
}

func serverCmd() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Start service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "configName",
				Aliases:  []string{"cn"},
				Usage:    "配置文件名称",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "configPath",
				Aliases:  []string{"cp"},
				Usage:    "配置文件搜索路径, 多个路径用逗号隔开",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "configType",
				Aliases:  []string{"ct"},
				Usage:    "配置文件类型, 默认yml",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return boot.Run(boot.RunConfig{
				ConfigName:       c.String("configName"),
				ConfigSearchPath: c.String("configPath"),
				ConfigType:       c.String("configType"),
			})
		},
	}
}
