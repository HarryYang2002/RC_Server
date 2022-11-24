package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	authpb "server/auth/api/gen/v1"
	"server/auth/auth"
	"server/auth/auth/wechat"
	"server/auth/dao"
	"server/auth/token"
	"server/shared/server"
	"time"
)

var addr = flag.String("addr", "8081", "address to listen")
var mongoURI = flag.String("mongo_uri", "your mongo uri", "mongo uri")
var privateKeyFile = flag.String("private_key_file", "auth/private.key", "private key file")
var wechatAppID = flag.String("wechat_app_id", "your APPID", "wechat app id")
var wechatAppSecret = flag.String("wechat_app_secret", "your APP Secret", "wechat app secret")

func main() {
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal("cannot connect database", zap.Error(err))
	}

	pkFile, err := os.Open(*privateKeyFile)
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}

	pkBytes, err := io.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "auth",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			//TODO 配置化Appid和AppSecret
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIDResolver: &wechat.Service{
					AppID:     *wechatAppID,
					AppSecret: *wechatAppSecret,
				},
				Mongo:          dao.NewMongo(mongoClient.Database("SZTURC")),
				Logger:         logger,
				TokenExpire:    2 * time.Hour,
				TokenGenerator: token.NewJWTTokenGen("SZTURC/server/auth", privKey),
			})
		},
		Addr: *addr,
	}))
}
