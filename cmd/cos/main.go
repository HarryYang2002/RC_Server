package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
)

func main() {
	u, _ := url.Parse("your url")
	b := &cos.BaseURL{BucketURL: u}
	ak := "your SecretID"
	sk := "your SecretKey"

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: ak,
			SecretKey: sk,
		},
	})
	name := "aaa.jpg"
	presignedURL, err := client.Object.GetPresignedURL(context.Background(), http.MethodPut, name, ak, sk, 1*time.Hour, nil)
	if err != nil {
		panic(any(err))
	}
	fmt.Println(presignedURL)
}
