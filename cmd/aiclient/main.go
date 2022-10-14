package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	coolenvpb "server/shared/coolenv"
)

func main() {
	conn, err := grpc.Dial("localhost:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(any(err))
	}

	ac := coolenvpb.NewAIServiceClient(conn)
	c := context.Background()
	res, err := ac.MeasureDistance(c, &coolenvpb.MeasureDistanceRequest{
		From: &coolenvpb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		To: &coolenvpb.Location{
			Latitude:  31,
			Longitude: 121,
		},
	})
	if err != nil {
		panic(any(err))
	}
	fmt.Printf("%+v\n", res)
}
