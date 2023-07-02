package main

import (
	"context"
	"fmt"
	"goMicroCode/message"
	"time"

	"go-micro.dev/v4"
)

func main() {

	service := micro.NewService(
		micro.Name("student.client"),
	)

	service.Init()

	studentService := message.NewStudentService("student_service", service.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
	time.Sleep(50 * time.Second)
}
