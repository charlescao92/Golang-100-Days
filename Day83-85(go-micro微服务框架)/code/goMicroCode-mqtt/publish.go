package main

import (
	"encoding/json"
	"goMicroCode-mqtt/message"
	"log"
	"time"

	"github.com/go-micro/plugins/v4/broker/mqtt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

// type Student struct {
// 	Name    string `json:"user_name"`
// 	Classes string `json:"user_classes"`
// 	Grade   int8   `json:"user_grade"`
// 	Phone   string `json:"user_phone"`
// }

func main() {

	service := micro.NewService(
		micro.Name("mqtt_server"),
		micro.Version("v1.0.0"),
		micro.Broker(mqtt.NewBroker()),
	)

	service.Init()

	brok := service.Server().Options().Broker
	if err := brok.Connect(); err != nil {
		log.Fatal(" broker connection failed, error : ", err.Error())
	}

	defer brok.Disconnect()

	for {
		student := &message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80, Phone: "12345678901"}
		msgBody, err := json.Marshal(student)
		if err != nil {
			log.Fatal(err.Error())
		}

		msg := &broker.Message{
			Header: map[string]string{
				"name": student.Name,
			},
			Body: msgBody,
		}

		err = brok.Publish("go.micro.srv.message", msg)
		if err != nil {
			log.Fatal(" 消息发布失败：%s\n", err.Error())
		} else {
			log.Print("消息发布成功")
		}

		time.Sleep(5 * time.Second)
	}

}
