package main

import (
	"fmt"
	"gitee.com/langya_platform/pkg/xzap"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	configName = "langya_platform.conf.toml"
	workSpace  = "./config"
)

func main() {
	envInit()
	signalInit()
}

func envInit() {
	currentEnv := os.Getenv("CUSTOM_RUNTIME_ENV")
	if currentEnv == "" {
		currentEnv = "dev"
	}
	cfgName := strings.Join([]string{workSpace, currentEnv, configName}, "/")
	err := readConfigFile(cfgName)
	if err != nil {
		panic(any(fmt.Errorf("Fatal error config file: %s \n", err)))
	}

	err = xzap.InitZLog(viper.GetStringSlice("log.outputPaths"),
		zapcore.Level(viper.GetInt("log.level")))
	if err != nil {
		panic(any(err))
	}
	xzap.Info("日志log", zap.Any("name", "申明辉"))
}

func readConfigFile(cfgName string) error {
	viper.SetConfigFile(cfgName)
	viper.SetConfigType("toml")
	return viper.ReadInConfig()
}

func signalInit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for sig := range ch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			log.Println("step 5: server exit success by SIGTERM/SIGINT/SIGUSR1")
			os.Exit(0)
		default:
			log.Println("step 5: unknown signal", sig)
		}
	}
}
