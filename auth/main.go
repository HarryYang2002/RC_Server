package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	authpb "server/auth/api/gen/v1"
	"server/auth/auth"
	"server/auth/auth/wechat"
	"server/auth/dao"
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

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/SZTURC?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect database", zap.Error(err))
	}

	s := grpc.NewServer()
	//TODO 配置化Appid和AppSecret
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wxe4f040053ecc73d0",
			AppSecret: "your APPSecret",
		},
		Mongo:  dao.NewMongo(mongoClient.Database("SZTURC")),
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
