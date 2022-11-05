package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	authpb "server/auth/api/gen/v1"
	carpb "server/car/api/gen/v1"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/server"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true,
			OrigName:    true,
		},
	))

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "trip",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
		{
			name:         "car",
			addr:         "localhost:8084",
			registerFunc: carpb.RegisterCarServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			logger.Sugar().Fatal("cannot register service %s: %v", s.name, err)
		}
	}
	addr := ":8080"
	logger.Sugar().Infof("grpc gateway started at %s", addr)
	logger.Sugar().Fatal(http.ListenAndServe(":8080", mux))
}
