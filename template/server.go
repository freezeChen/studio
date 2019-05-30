/*
   @Time : 2019-05-08 17:09
   @Author : frozenchen
   @File : server
   @Software: studio
*/
package template

var HttpServer_template = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : server
   @Software: {{.Appname}}
*/
package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{.Package}}/service"
	"{{.Package}}/proto"
)

var (
	svc *service.Service
)

func New(s *service.Service) (engine *gin.Engine) {
	engine = gin.Default()
	svc = s
	initRouter(engine)
	return
}

func initRouter(e *gin.Engine) {
	g := e.Group("/demo")
	{
		g.GET("/hello", hello)
	}
}

func hello(ctx *gin.Context) {
	reply,_ :=svc.HelloService.Hello(ctx,&proto.Req{S:"studio"})
	ctx.JSON(http.StatusOK, reply.Message)
}
`
