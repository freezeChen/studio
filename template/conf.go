/*
   @Time : 2019-05-09 16:41
   @Author : frozenchen
   @File : conf
   @Software: studio
*/
package template

var ConfDemo_template = `
package conf

import (
	"github.com/freezeChen/studio-library/conf"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/util"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-config"
)

type Config struct {
	Name    string
	Version string
	Env     string
	Debug   bool
	Log		*zlog.Config
	Mysql   *mysql.Config
	Redis   *redis.Config
}

func Init() (*Config, error) {
	var (
		Conf = &Config{}
	)
	cfg := config.NewConfig()

	source := conf.LoadFileSource(util.GetCurrentDirectory() + "/conf.yaml")
	cfg.Load(source)
	if err := cfg.Scan(Conf); err != nil {
		return nil, err
	}

	return Conf, nil
}

`

var ConfFile_template = `
name: "{{.Appname}}"
env: "dev"
version: "v0.0.1"
Debug: true
mysql:
  source: "test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"
  active: 100
  idle: 20
log:
  name: "{{.Appname}}"
  debug: true
  writeKafka: false
  kafkaAddr: "127.0.0.1:9092"
redis:
  addr: "127.0.0.1:6379"
  auth: ""
  active: 100
  idle: 20
  dialTimeout: "1ms"
  readTimeout: "1ms"
  writeTimeout: "1ms"
  idleTimeout: "1ms"
kafka: "127.0.0.1:9092"
`
