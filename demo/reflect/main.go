package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 1
	//getType(n)
	getValue(n)
	setValue(&n)
	fmt.Println(n)
}

func getType(str interface{}) {
	t := reflect.TypeOf(str)
	fmt.Println(t)
	fmt.Printf("%T", t)
}
func getValue(str interface{}) {
	v := reflect.ValueOf(str)
	k := v.Kind()
	switch k {
	case reflect.Int:
		ret := int(v.Int())
		fmt.Printf("%v,%T", ret, ret)
	}
}

func setValue(str interface{}) {
	v := reflect.ValueOf(str)
	t := v.Elem().Kind()
	switch t {
	case reflect.Int:
		v.Elem().SetInt(100)
	}
}
