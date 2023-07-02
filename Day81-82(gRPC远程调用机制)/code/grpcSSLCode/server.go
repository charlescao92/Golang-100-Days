package main

import (
	"context"
	"fmt"
	"grpcSSLCode/message"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println(" 服务端 Add方法 ")
	result := request.Args1 + request.Args2
	fmt.Println(" 计算结果是：", result)
	response = new(message.Response)
	response.Code = 1
	response.Message = "执行成功"
	return response, nil
}

func main() {

	// 方法1
	// cert, err := tls.LoadX509KeyPair("./keys/server.crt", "./keys/server.key")
	// if err != nil {
	// 	grpclog.Fatal("加载在证书文件失败", err)
	// }
	// creds := credentials.NewTLS(&tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// })

	// 方法2
	creds, err := credentials.NewServerTLSFromFile("./keys/server.crt", "./keys/server.key")
	if err != nil {
		grpclog.Fatal("加载在证书文件失败", err)
	}

	//实例化grpc server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
