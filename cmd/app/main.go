package main

import (
	"context"
	"fmt"
	app2 "gitee.com/langya_platform/langya/platform/app"
	"gitee.com/langya_platform/pkg/xzap"
	"gitee.com/langya_platform/service/app"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
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
	go startAppServ()
	signalInit()
}

func startAppServ() {
	// grpc服务
	err := app.NewPlatformAppService(context.Background())
	if err != nil {
		panic(any("failed to init PlatformAppService"))
	}

	// http服务
	conn, err := grpc.Dial(
		"localhost:9901",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	lamux := runtime.NewServeMux()
	// Register Greeter
	err = app2.RegisterLangYaPlatformHandler(context.Background(), lamux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	lpAppServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 9900),
		Handler: lamux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0" + fmt.Sprintf(":%d", 9900))
	log.Fatalln(lpAppServer.ListenAndServe())
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
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM) // signal包将输入信号转发到ch,如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号
	for sig := range ch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			log.Println("step 5: server exit success by SIGTERM/SIGINT")
			os.Exit(0)
		default:
			log.Println("step 5: unknown signal", sig)
		}
	}
}
