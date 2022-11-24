package main

import (
	"context"
	"github.com/namsral/flag"
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

var addr = flag.String("addr", ":8083", "address to listen")
var mongoURI = flag.String("mongo_uri", "your mongo uri", "mongo uri")
var cosAddr = flag.String("cos_addr", "your cos addr", "cos address")
var cosSecID = flag.String("cos_sec_id", "your cos secret ID", "cos secret id")
var cosSecKey = flag.String("cos_sec_key", "your cos secret key", "cos secret key")

func main() {
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	db := mongoClient.Database("SZTURC")

	st, err := cos.NewService(*cosAddr, *cosSecID, *cosSecKey)
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   *addr,
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
