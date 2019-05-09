/*
   @Time : 2019-05-09 13:42
   @Author : frozenchen
   @File : makefile
   @Software: studio
*/
package template

var Makefile_template = `build:
	go build -o {{.Appname}}

`
