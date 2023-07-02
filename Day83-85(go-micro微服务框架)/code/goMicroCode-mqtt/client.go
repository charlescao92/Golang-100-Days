package main

import (
	"encoding/json"
	"fmt"
	"goMicroCode-mqtt/message"
	"log"
	"time"

	"github.com/go-micro/plugins/v4/broker/mqtt"
	_ "github.com/go-micro/plugins/v4/broker/mqtt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

func main() {

	service := micro.NewService(
		micro.Name("mqtt_client"),
		micro.Version("v1.0.0"),
		micro.Broker(mqtt.NewBroker()),
	)

	service.Init()

	pubSub := service.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatal(" broker connection failed, error : ", err.Error())
	}

	subscriber, err := pubSub.Subscribe("go.micro.srv.message", func(event broker.Event) error {
		var req *message.StudentRequest
		if err := json.Unmarshal(event.Message().Body, &req); err != nil {
			fmt.Println("json.Unmarshal err", err.Error())
			return err
		}
		log.Print(" 接收到信息：", req)
		return nil
	})

	if err != nil {
		fmt.Println(subscriber.Topic(), "主题订阅失败！ ", err.Error())
	} else {
		fmt.Println(subscriber.Topic(), "主题订阅成功！")
	}

	defer func() {
		fmt.Println("client close conn and Unsubscribe")
		pubSub.Disconnect()      //关闭链接
		subscriber.Unsubscribe() //取消订阅
	}()

	time.Sleep(30 * time.Second)
}
