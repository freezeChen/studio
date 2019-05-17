/*
   @Time : 2019-05-08 17:08
   @Author : frozenchen
   @File : main
   @Software: studio
*/
package template

var Main_web = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : main
   @Software: {{.Appname}}
*/
package main

import (
	"github.com/micro/go-web"
	"github.com/micro/go-micro/client"
	"{{.Appname}}/conf"
	"{{.Appname}}/server/http"
	"{{.Appname}}/proto"
	"{{.Appname}}/service"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	svc := web.NewService(
	web.Name("go.micro.web.hello"),
	web.Address(":8080"))
	svc.Init()
    helloService := proto.NewHelloService("go.micro.srv.hello", client.DefaultClient)
	s:=service.New(cfg)
	s.HelloService = helloService
	engine := http.New(s)
	svc.Handle("/", engine)
	if err := svc.Run(); err != nil {
		return
	}
}
`

var Main_srv = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : main
   @Software: {{.Appname}}
*/
package main

import (
	"github.com/micro/go-micro"
	"{{.Appname}}/conf"
	"{{.Appname}}/proto"
	"{{.Appname}}/service"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	svc := micro.NewService(
	micro.Name("go.micro.srv.hello"),
	micro.Address(":8081"))
	svc.Init()
	s := service.New(cfg)
	err = proto.RegisterHelloHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
`
