package app

import (
	"context"
	"fmt"
	"gitee.com/langya_platform/langya/platform/app"
	"gitee.com/langya_platform/pkg/xzap"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewPlatformAppService 注册grpc服务
func NewPlatformAppService(ctx context.Context) error {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	langyaAppSrv := new(LangyaPlatformService)
	// 将服务描述(server)及其具体实现(LangYaPlatformServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
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
