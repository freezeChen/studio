/*
   @Time : 2019-05-08 16:43
   @Author : frozenchen
   @File : service
   @Software: studio
*/
package service

import "{{cookiecutter.app_name}}/dao"

type Service struct {
	dao *dao.Dao
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}
	return s
}
