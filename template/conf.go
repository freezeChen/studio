/*
   @Time : 2019-05-09 16:41
   @Author : frozenchen
   @File : conf
   @Software: studio
*/
package template

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"os"
	"path/filepath"
)

var ConfDemo_template = `
package conf

type Config struct {
	Name string
	Version string
	Env string
	Mysql 
}

func Init() {

}

`
var (
	Conf = &Config{}
)

type Config struct {
	Name    string
	Version string
	Env     string
	Mysql   *mysqlConfig
	Redis   *redisConfig
}

type mysqlConfig struct {
	Source string
	Active int
	Idle   int
}

type redisConfig struct {
	Addr   string
	Auth   string
	Active int
	Idle   int
}

func Init() {
	cfg := config.NewConfig()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	source := file.NewSource(file.WithPath(dir + "/conf.yaml"))

	err := cfg.Load(source)

	cfg.Scan(Conf)
	if err != nil {
		fmt.Println(err)
	}

}
