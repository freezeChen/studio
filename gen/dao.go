package testdata

import (
	"context"
	"github.com/freezeChen/studio-library/redis"

)

// Demo test struct
type Demo struct {
	ID    int64
	Title string
}

// Dao .
type Dao struct {
	cache *redis.Redis
}

// New .
func New() *Dao {
	return &Dao{}
}

//go:generate kratos tool genbts
type _bts interface {
	// bts: -batch=2 -max_group=20 -batch_err=break -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
	Demos(c context.Context, keys []int64) (map[int64]*Demo, error)
	// bts: -batch=2 -max_group=20 -batch_err=continue -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
	Demos1(c context.Context, keys []int64) (map[int64]*Demo, error)
	// bts: -sync=true -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
	Demo(c context.Context, key int64) (*Demo, error)
	// bts: -paging=true
	Demo1(c context.Context, key int64, pn int, ps int) (*Demo, error)
	// bts: -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
	None(c context.Context) (*Demo, error)
}
