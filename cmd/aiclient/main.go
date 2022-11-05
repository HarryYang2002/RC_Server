package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"server/car/mq/amqpclt"
	coolenvpb "server/shared/coolenv"
	"server/shared/server"
	"time"
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

	idRes, err := ac.LicIdentity(c, &coolenvpb.IdentityRequest{
		Photo: []byte{1, 2, 3, 4, 5},
	})
	if err != nil {
		panic(any(err))
	}
	fmt.Printf("%+v\n", idRes)

	_, err = ac.SimulateCarPos(c, &coolenvpb.SimulateCarPosRequest{
		CarId: "car123",
		InitialPos: &coolenvpb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		//Type: coolenvpb.PosType_RANDOM, //随机数字
		Type: coolenvpb.PosType_NINGBO,
	})
	if err != nil {
		panic(any(err))
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		panic(any(err))
	}

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(any(err))
	}

	sub, err := amqpclt.NewSubscriber(amqpConn, "pos_sim", logger)
	if err != nil {
		panic(any(err))
	}

	ch, cleanUp, err := sub.SubscribeRaw(c)
	defer cleanUp()

	if err != nil {
		panic(any(err))
	}

	tm := time.After(10 * time.Second)
	for {
		shouldStop := false
		select {
		case msg := <-ch:
			fmt.Printf("%s\n", msg.Body)
			//var update coolenvpb.CarPosUpdate
			//err = json.Unmarshal(msg.Body, &update)
			//if err != nil {
			//	panic(err)
			//}
			//fmt.Printf("%+v\n", &update)
		case <-tm:
			shouldStop = true
		}
		if shouldStop {
			break
		}
	}

	_, err = ac.EndSimulateCarPos(c, &coolenvpb.EndSimulateCarPosRequest{
		CarId: "car123",
	})
	if err != nil {
		panic(any(err))
	}

}
