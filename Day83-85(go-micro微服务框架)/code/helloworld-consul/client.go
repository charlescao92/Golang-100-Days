package main

import (
	"context"
	"fmt"
	"time"

	pb "helloworld-consul/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
)

func main() {
	registry := consul.NewRegistry()

	service := micro.NewService(
		//micro.Name("helloworld-1"),
		micro.Registry(registry),
	)

	service.Init()

	greeterService := pb.NewGreeterService("helloworld-consul", service.Client())

	res, err := greeterService.Hello(context.TODO(), &pb.Request{Name: "charles"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Greeting)
	time.Sleep(50 * time.Second)
}
