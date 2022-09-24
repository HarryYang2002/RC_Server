package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	trippb "server/protoc/gen/go"
	trip "server/tripservice"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true,
			OrigName:    true,
		}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8081",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}
}
