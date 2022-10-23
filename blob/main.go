package main

import (
	"context"
	"log"
	blobpb "server/blob/api/gen/v1"
	"server/blob/blob"
	"server/blob/cos"
	"server/blob/dao"
	"server/shared/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	st, err := cos.NewService(
		"your address",
		"your sid",
		"your skey")
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   ":8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Storage: st,
				Mongo:   dao.NewMongo(db),
				Logger:  logger,
			})
		},
	}))
}
