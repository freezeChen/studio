/*
   @Time : 2019-05-08 17:09
   @Author : frozenchen
   @File : dao
   @Software: studio
*/
package template

var Dao_template = `/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : dao
   @Software: {{.Appname}}
*/
package dao

import (
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"{{.Package}}/conf"
)

type Dao struct {
	Db    xorm.EngineInterface
	Redis *redis.Redis
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		Db:    mysql.New(c.Mysql),
		Redis: redis.New(c.Redis),
	}
	return
}
`
