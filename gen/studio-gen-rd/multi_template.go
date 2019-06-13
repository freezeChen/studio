package main

import (
	"strings"
)

var _multiGetTemplate = `
// NAME {{or .Comment "get data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, ids []KEY {{.ExtraArgsType}}) (res map[KEY]VALUE, err error) {
    l := len(ids)
	if l == 0 {
		return
	}
	{{if .EnableBatch}}
		mutex := sync.Mutex{}
		for i:=0;i < l; i+= GROUPSIZE * MAXGROUP {
			var subKeys []KEY
			{{if .BatchErrBreak}}
				group := errgroup.WithContext(c)
			{{else}}
				group := &errgroup.Group{}
				
			{{end}}
			if (i + GROUPSIZE * MAXGROUP) > l {
				subKeys = ids[i:]
			} else {
				subKeys = ids[i : i+GROUPSIZE * MAXGROUP]
			}
			subLen := len(subKeys)
			for j:=0; j< subLen; j += GROUPSIZE {
				var ks []KEY
				if (j+GROUPSIZE) > subLen {
					ks = subKeys[j:]
				} else {
					ks = subKeys[j:j+GROUPSIZE]
				}
				group.Go(func(ctx context.Context) (err error) {
					keysMap := make(map[string]KEY, len(ks))
					keys := make([]string, 0, len(ks))
					for _, id := range ks {
						key := {{.KeyMethod}}(id{{.ExtraArgs}})
						keysMap[key] = id
						keys = append(keys, key)
					}
					replies, err := d.mc.CacheGetMulti(c, keys)
					if err != nil {

						return
					}
					for _, key := range keys {
						{{if .GetSimpleValue}}
							var v string
							err = replies.Scan(key, &v)
						{{else}}
							var v VALUE
							{{if .GetDirectValue}}
								err = replies.Scan(key, &v)
							{{else}}
								{{if .PointType}}
									v = &{{.OriginValueType}}{}
									err = replies.Scan(key, v)
								{{else}}
									v = {{.OriginValueType}}{}
									err = replies.Scan(key, &v)
								{{end}}
							{{end}}
						{{end}}
						if err != nil {
							return
						}
						{{if .GetSimpleValue}}
							r, err := {{.ConvertBytes2Value}}
							if err != nil {
								return res, err
							}
							mutex.Lock()
							if res == nil {
								res = make(map[KEY]VALUE, len(keys))
							}
							res[keysMap[key]] = {{.ValueType}}(r)
							mutex.Unlock()
						{{else}}
							mutex.Lock()
							if res == nil {
								res = make(map[KEY]VALUE, len(keys))
							}
							res[keysMap[key]] = v
							mutex.Unlock()
						{{end}}
					}
				return
				})
			}
			err1 := group.Wait()
			if err1 != nil {
				err = err1
			{{if .BatchErrBreak}}
				break
			{{end}}
			}
		}
	{{else}}
		keysMap := make(map[string]KEY, l)
		keys := make([]string, 0, l)
		for _, id := range ids {
			key := {{.KeyMethod}}(id{{.ExtraArgs}})
			keysMap[key] = id
			keys = append(keys, key)
		}
		replies, err := d.mc.CacheGetMulti(c, keys)
		if err != nil {
			return
		}
		for _, key := range keys {
			{{if .GetSimpleValue}}
				var v string
				err = replies.Scan(key, &v)
			{{else}}
				{{if .PointType}}
					v := &{{.OriginValueType}}{}
					err = replies.Scan(key, v)
				{{else}}
					v := {{.OriginValueType}}{}
					err = replies.Scan(key, &v)
				{{end}}
			{{end}}
			if err != nil {
				return
			}
			{{if .GetSimpleValue}}
				r, err := {{.ConvertBytes2Value}}
				if err != nil {
					return res, err
				}
				if res == nil {
					res = make(map[KEY]VALUE, len(keys))
				}
				res[keysMap[key]] = {{.ValueType}}(r)
			{{else}}
				if res == nil {
					res = make(map[KEY]VALUE, len(keys))
				}
				res[keysMap[key]] = v
			{{end}}
		}
	{{end}}
	return
}
`

var _multiSetTemplate = `
// NAME {{or .Comment "Set data to redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, values map[KEY]VALUE {{.ExtraArgsType}}) (err error) {
	if len(values) == 0 {
		return
	}
	for id, val := range values {
		key := {{.KeyMethod}}(id{{.ExtraArgs}})
		{{if .SimpleValue}}
			bs := {{.ConvertValue2Bytes}}
			item := &redis.Item{Key: key, Value: bs, Expiration: {{.ExpireCode}}}
		{{else}}
			item := &redis.Item{Key: key, Object: val, Expiration: {{.ExpireCode}}}
		{{end}}
		if err = d.mc.CaCheSet(c, item); err != nil {
			
			return
		}
	}
	return
}
`
var _multiAddTemplate = strings.Replace(_multiSetTemplate, "Set", "Add", -1)
var _multiReplaceTemplate = strings.Replace(_multiSetTemplate, "Set", "Replace", -1)

var _multiDelTemplate = `
// NAME {{or .Comment "delete data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, ids []KEY {{.ExtraArgsType}}) (err error) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		key := {{.KeyMethod}}(id{{.ExtraArgs}})
		if err = d.mc.Delete(c, key); err != nil {
			if err == redis.ErrNotFound {
				err = nil
				continue
			}
		
			return
		}
	}
	return
}
`
