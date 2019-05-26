package main

import (
	"flag"
	"fmt"
	"time"
)


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

var AppName = flag.String("name", "", "app name")
var ProjectName =flag.String("project", "", "project name")
var Type = flag.String("type","web","gen web or srv")
var Single = flag.Bool("single",true,`when single is false;project path like project/name`)
func main() {
	flag.Parse()
var c config
	switch *Type {
	case "web":

		c =config{
			Appname: *AppName,
			Type:    *Type,
			Project: *ProjectName,
			Single:  *Single,
			Author:  "freezeChen",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Files: []file{
				{"file22.go",Service_web_template},
			},
		}
	case "srv":
		c =config{
			Appname: *AppName,
			Type:    *Type,
			Project: *ProjectName,
			Single:  *Single,
			Author:  "freezeChen",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Files: []file{
				{"file22.go",Service_web_template},
			},
		}

	}





	if err := create(c); err != nil {
		fmt.Println(err.Error())
		return
	}

}
