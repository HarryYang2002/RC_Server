package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/trip"
	"server/rental/trip/client/car"
	"server/rental/trip/client/poi"
	"server/rental/trip/client/profile"
	"server/rental/trip/dao"
	"server/shared/server"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/SZTURC?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect database", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Logger:            logger,
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager:    &car.Manager{},
				ProfileManage: &profile.Manager{},
				PoiManager:    &poi.Manager{},
				Mongo:         dao.NewMongo(mongoClient.Database("SZTURC")),
				Logger:        logger,
			})
		},
		Addr: ":8082",
	}))
}
