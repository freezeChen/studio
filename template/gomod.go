/*
   @Time : 2019-05-09 13:41
   @Author : frozenchen
   @File : gomod
   @Software: studio
*/
package template


var Go_mod = `module {{.Appname}}

go 1.12

require (
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)


`
