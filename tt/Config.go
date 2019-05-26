package main

import (
	"ttt/util"
)

type config struct {
	Appname string
	Type    string //web,srv
	Project string
	Dir     string
	Single  bool
	Author  string
	Time    string
	Files   []file

}

func (c config) getPath()string  {
	if c.Single {
		return util.GetCurrentDirectory()+"/"+c.Appname
	}else {
		return util.GetCurrentDirectory()+"/"+c.Project
	}
}