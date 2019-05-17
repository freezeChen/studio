/*
   @Time : 2019-05-08 17:09
   @Author : frozenchen
   @File : service
   @Software: studio
*/
package template

var Service_template = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : service
   @Software: {{.Appname}}
*/
package service

import (
	"{{.Appname}}/conf"
	"{{.Appname}}/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}
	return s
}

`

var Service_web_template = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : service
   @Software: {{.Appname}}
*/
package service

import (
	"{{.Appname}}/conf"
	"{{.Appname}}/dao"
	"{{.Appname}}/proto"
)

type Service struct {
	dao *dao.Dao
	HelloService proto.HelloService
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}
	return s
}


`



var Service_handler_teplate=`/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : service
   @Software: {{.Appname}}
*/
package service

import (
	"context"
	"fmt"
	"{{.Appname}}/proto"
)

func (Service) Hello(ctx context.Context, req *proto.Req, reply *proto.Reply) error {
	reply.Message = fmt.Sprintf("hello %s, Congratulations you success call rpc service!", req.S)
	return nil
}

`
