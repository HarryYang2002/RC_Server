package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/trip"
)

func main() {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
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
