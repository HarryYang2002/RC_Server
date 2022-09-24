package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	trippb "server/protoc/gen/go"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect server: %v",err)
	}
	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})
	if err != nil {
		log.Fatalf("cannto call GetTrip: %v",err)
	}

	fmt.Println(r)
}
