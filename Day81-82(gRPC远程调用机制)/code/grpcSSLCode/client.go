package main

import (
	"context"
	"fmt"
	"grpcSSLCode/message"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./keys/server.crt", "")
	if err != nil {
		panic(err.Error())
	}

	grpc.WithInsecure()

	// 创建 gRPC 连接，指定服务器的地址和凭据
	conn, err := grpc.Dial("localhost:8092", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)

	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	fmt.Println(response.GetCode(), response.GetMessage())
}
