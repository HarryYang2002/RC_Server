package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"server/rental/ai"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/profile"
	profiledao "server/rental/profile/dao"
	"server/rental/trip"
	"server/rental/trip/client/car"
	"server/rental/trip/client/poi"
	profClient "server/rental/trip/client/profile"
	tripdao "server/rental/trip/dao"
	coolenvpb "server/shared/coolenv"
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

	ac, err := grpc.Dial("localhost:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect aiService", zap.Error(err))
	}

	db := mongoClient.Database("SZTURC")

	profService := &profile.Service{
		Mongo:  profiledao.NewMongo(db),
		Logger: logger,
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Logger:            logger,
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{},
				ProfileManage: &profClient.Manager{
					Fetcher: profService,
				},
				PoiManager: &poi.Manager{},
				DistanceCalc: &ai.Client{
					AIClient: coolenvpb.NewAIServiceClient(ac),
				},
				Mongo:  tripdao.NewMongo(db),
				Logger: logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profService)
		},
		Addr: ":8082",
	}))
}
