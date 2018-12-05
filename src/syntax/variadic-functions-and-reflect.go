package main

import (
	"fmt"
	"reflect"
)

func variadicReflect(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main() {
	variadicReflect(9527, "xiucai", true, 13.14, []string{"huawen", "huawu"}, map[string]int{"qiuxiang": 250})
}
