/*
   @Time : 2019-05-09 13:41
   @Author : frozenchen
   @File : gomod
   @Software: studio
*/
package template


var Go_mod = `module {{.Package}}

go 1.12

require (
	github.com/gin-gonic/gin v1.3.0
	github.com/micro/go-micro v1.7.0
	github.com/freezeChen/studio-library v0.0.16
)

`
