package boot

import (
	"os"
	"os/signal"
	"strings"

	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/config"
	"github.com/alibaba/ioc-golang/logger"
)

type RunConfig struct {
	// default: config
	ConfigName string
	// Search path of config files
	//
	// default: ., ./config, ./configs;
	//
	// priority: ./ > ./config > ./configs
	// 多个path使用逗号隔开
	ConfigSearchPath string
	// default: yml
	ConfigType string
}

// Run 运行
func Run(con RunConfig) error {
	logger.Disable()
	printLogo()

	var opts []config.Option
	if con.ConfigName != "" {
		opts = append(opts, config.WithConfigName(con.ConfigName))
	}
	if con.ConfigType == "" {
		con.ConfigType = "yml"
	}
	opts = append(opts, config.WithConfigType(con.ConfigType))
	if con.ConfigSearchPath != "" {
		searchPathArr := strings.Split(con.ConfigSearchPath, ",")
		opts = append(opts, config.WithSearchPath(searchPathArr...))
	}

	if err := ioc.Load(opts...); err != nil {
		return err
	}

	configurer, err := GetModuleConfigurerSingleton()
	if err != nil {
		return err
	}
	clean, err := configurer.Configure()
	if err != nil {
		return nil
	}

	listeningSignal(func() {
		for _, fn := range clean {
			fn()
		}
	})
	return nil
}

func listeningSignal(doExit func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	doExit()
}
