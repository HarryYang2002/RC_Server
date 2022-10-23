package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	blobpb "server/blob/api/gen/v1"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(any(err))
	}
	c := blobpb.NewBlobServiceClient(conn)

	ctx := context.Background()
	//res, err := c.CreateBlob(ctx, &blobpb.CreateBlobRequest{
	//	AccountId:           "account_2",
	//	UploadUrlTimeoutSec: 1000,
	//})
	//res, err := c.GetBlob(ctx, &blobpb.GetBlobRequest{
	//	Id: "63557570524362c132578ef1",
	//})
	res, err := c.GetBlobURL(ctx, &blobpb.GetBlobURLRequest{
		Id:         "63557570524362c132578ef1",
		TimeoutSec: 100,
	})
	if err != nil {
		panic(any(err))
	}

	fmt.Printf("%+v\n", res)
}
