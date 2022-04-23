package app

import (
	"context"
	"fmt"
	"gitee.com/langya_platform/langya/platform/app"
	"gitee.com/langya_platform/pkg/xzap"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewPlatformAppService 注册grpc服务
func NewPlatformAppService(ctx context.Context) error {
	server := grpc.NewServer()
	langyaAppSrv := new(LangyaPlatformService)
	app.RegisterLangYaPlatformServer(server, langyaAppSrv)

	// Serve gRPC Server
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt32("grpc.port")))
	if err != nil {
		return err
	}
	xzap.Info("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", viper.GetInt32("grpc.port")))
	go func() {
		if err := server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return nil
}
