package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	zs := Student{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	t := reflect.TypeOf(zs)
	fmt.Printf("name:%v,kind:%v\n", t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%v,tag:%v,type:%v\n", field.Name, field.Tag, field.Type)
	}

	name, ok := t.FieldByName("Name")
	if ok {
		fmt.Println(name.Tag, name.Name, name.Type)
	}

	v := reflect.ValueOf(zs)
	nameValue := v.FieldByName("Name").String()
	ageValue := v.FieldByName("Age").Int()
	fmt.Printf("name:%v,age:%v\n", nameValue, ageValue)
}
