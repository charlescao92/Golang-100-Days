package main

import (
	"context"
	"fmt"
	"time"

	"go-micro.dev/v4"
	pb "helloworld/proto"
)

func main() {

	service := micro.NewService(
		micro.Name("helloworld"),
	)

	service.Init()

	greeterService := pb.NewGreeterService("helloworld", service.Client())

	res, err := greeterService.Hello(context.TODO(), &pb.Request{Name: "charles"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Greeting)
	time.Sleep(50 * time.Second)
}
