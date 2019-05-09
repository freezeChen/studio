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

import "{{.Appname}}/dao"

type Service struct {
	dao *dao.Dao
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}
	return s
}

`
