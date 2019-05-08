/*
   @Time : 2019-05-08 16:42
   @Author : frozenchen
   @File : http
   @Software: studio
*/
package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{cookiecutter.app_name}}/service"
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
	ctx.JSON(http.StatusOK, "hello")
}
