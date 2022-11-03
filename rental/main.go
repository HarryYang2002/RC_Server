package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	blobpb "server/blob/api/gen/v1"
	carpb "server/car/api/gen/v1"
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
	"time"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/SZTURC?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	db := mongoClient.Database("SZTURC")

	blobConn, err := grpc.Dial("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect blob service", zap.Error(err))
	}

	ac, err := grpc.Dial("localhost:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect aiservice", zap.Error(err))
	}
	aiClient := &ai.Client{
		AIClient:  coolenvpb.NewAIServiceClient(ac),
		UseRealAI: false,
	}

	profService := &profile.Service{
		BlobClient:        blobpb.NewBlobServiceClient(blobConn),
		PhotoGetExpire:    5 * time.Second,
		PhotoUploadExpire: 10 * time.Second,
		IdentityResolver:  aiClient,
		Mongo:             profiledao.NewMongo(db),
		Logger:            logger,
	}

	carConn, err := grpc.Dial("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect car service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{
					CarService: carpb.NewCarServiceClient(carConn),
				},
				ProfileManage: &profClient.Manager{
					Fetcher: profService,
				},
				PoiManager:   &poi.Manager{},
				DistanceCalc: aiClient,
				Mongo:        tripdao.NewMongo(db),
				Logger:       logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profService)
		},
	}))
}
