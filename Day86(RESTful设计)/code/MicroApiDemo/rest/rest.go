package main

import (
	"log"

	restful "github.com/emicklei/go-restful/v3"

	"MicroApiDemo/proto"
	"context"

	"fmt"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/web"
)

type Student struct {
}

var (
	cli proto.StudentService
)

func (s *Student) GetStudent(req *restful.Request, rsp *restful.Response) {

	name := req.PathParameter("name")
	fmt.Println(name)
	response, err := cli.GetStudent(context.TODO(), &proto.Request{
		Name: name,
	})

	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.student"),
	)

	service.Init()

	cli = proto.NewStudentService("go.micro.srv.student", client.DefaultClient)

	student := new(Student)
	ws := new(restful.WebService)
	ws.Path("/student")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{name}").To(student.GetStudent))

	wc := restful.NewContainer()
	wc.Add(ws)

	service.Handle("/", wc)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
