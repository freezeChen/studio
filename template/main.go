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
	"github.com/micro/go-micro/web"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/client"
	"{{.Package}}/conf"
	"{{.Package}}/server/http"
	"{{.Package}}/proto"
	"{{.Package}}/service"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	zlog.InitLogger(cfg.Log)
	svc := web.NewService(
	web.Name("go.micro.web.hello"),
	web.Address(":8080"),
	web.RegisterInterval(20*time.Second),
	web.RegisterTTL(30*time.Second))

	if err := svc.Init(); err != nil {
		panic(err)
	}

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
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro"
	"{{.Package}}/conf"
	"{{.Package}}/proto"
	"{{.Package}}/service"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	zlog.InitLogger(cfg.Log)
	svc := micro.NewService(
	micro.Name("go.micro.srv.hello"),
	micro.Address(":8081"),
	micro.RegisterInterval(20*time.Second),
	micro.RegisterTTL(30*time.Second))

	svc.Init()
	s := service.New(cfg)
	err = proto.RegisterHelloHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
`
