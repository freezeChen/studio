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
	"github.com/micro/go-micro"
	"{{.Appname}}/server/http"
	"{{.Appname}}/service"
)

func main() {
	svc := micro.NewService()
	svc.Init()

	http.New(service.New())
	if err := svc.Run(); err != nil {
		return
	}
}
`

var Main_srv = `

`
