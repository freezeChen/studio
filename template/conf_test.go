/*
   @Time : 2019-05-09 17:28
   @Author : frozenchen
   @File : conf_test.go
   @Software: studio
*/
package template

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	config, e := Init()
	if e != nil {
		t.Error(e)
	}

	fmt.Println(fmt.Sprintf("%+v", config))

}
