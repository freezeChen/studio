/*
   @Time : 2019-05-08 17:09
   @Author : frozenchen
   @File : dao
   @Software: studio
*/
package template

var Dao_template =
`/*
   @Time : {{.Time}}
   @Author : {{.Author}}
   @File : dao
   @Software: {{.Appname}}
*/
package dao

type Dao struct {
}

func New() (dao *Dao) {

	dao = &Dao{

	}

	return
}

`
