package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	authpb "server/auth/api/gen/v1"
	"server/auth/auth"
	"server/auth/auth/wechat"
)

func main() {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	//TODO 配置化Appid和AppSecret
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID: "wxe4f040053ecc73d0",
			AppSecret: "your AppSecret",
		},
		Logger: logger,
	})

	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}

// zapLogger可以自定义
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
