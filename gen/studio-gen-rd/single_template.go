package main

import (
	"strings"
)

var _singleGetTemplate = `
// NAME {{or .Comment "get data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY {{.ExtraArgsType}}) (res VALUE, err error) {
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	{{if .GetSimpleValue}}
		var v string
		err = d.mc.CacheGet(c, key).Scan(&v)
	{{else}}
		{{if .GetDirectValue}}
			err = d.mc.CacheGet(c, key).Scan(&res)
		{{else}}
			{{if .PointType}}
				res = &{{.OriginValueType}}{}
				if err = d.mc.CacheGet(c, key).Scan(res); err != nil {
					res = nil
					return 
				}
			{{else}}
				err = d.mc.CacheGet(c, key).Scan(&res)
			{{end}}
		{{end}}
	{{end}}
	{{if .GetSimpleValue}}
		r, err := {{.ConvertBytes2Value}}
		if err != nil {
		
			return
		}
		res = {{.ValueType}}(r)
	{{end}}
	return
}
`

var _singleSetTemplate = `
// NAME {{or .Comment "Set data to redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY, val VALUE {{.ExtraArgsType}}) (err error) {
	{{if .PointType}}
      if val == nil {
        return 
      }
	{{end}}
	{{if .LenType}}
      if len(val) == 0 {
        return 
      }
	{{end}}
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	{{if .SimpleValue}}
		bs := {{.ConvertValue2Bytes}}
		item := &redis.Item{Key: key, Value: bs, Expiration: {{.ExpireCode}}}
	{{else}}
		item := &redis.Item{Key: key, Object: val, Expiration: {{.ExpireCode}}}
	{{end}}
	err = d.mc.CaCheSet(c,item)

	return
}
`
var _singleAddTemplate = strings.Replace(_singleSetTemplate, "Set", "Add", -1)
var _singleReplaceTemplate = strings.Replace(_singleSetTemplate, "Set", "Replace", -1)

var _singleDelTemplate = `
// NAME {{or .Comment "delete data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY {{.ExtraArgsType}}) (err error) {
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	err = d.mc.Delete(c,key)
	return
}
`
