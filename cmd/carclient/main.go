package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	carpb "server/car/api/gen/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(any(err))
	}
	cs := carpb.NewCarServiceClient(conn)
	c := context.Background()
	//for i := 0; i < 5; i++ {
	//	res, err := cs.CreateCar(c, &carpb.CreateCarRequest{})
	//	if err != nil {
	//		panic(any(err))
	//	}
	//	fmt.Printf("created car: %s\n",res.Id)
	//}
	res, err := cs.GetCars(c, &carpb.GetCarsRequest{})
	if err != nil {
		panic(any(err))
	}
	for _, car := range res.Cars {
		_, err := cs.UpdateCar(c, &carpb.UpdateCarRequest{
			Id:     car.Id,
			Status: carpb.CarStatus_LOCKED,
		})
		if err != nil {
			fmt.Printf("cannot reset cat %q: %v", car.Id, err)
		}
	}
	fmt.Printf("%d cars are resert", len(res.Cars))
}
