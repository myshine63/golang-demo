package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := student{
		name: "tom",
		Age:  10,
	}
	typeOfStu := reflect.TypeOf(stu)
	for i := 0; i < typeOfStu.NumField(); i++ {
		field := typeOfStu.Field(i)
		fmt.Printf("字段名字:%v,字段标签：%v,是否为匿名字段：%v\n", field.Name, field.Tag, field.Anonymous)
	}
}

type student struct {
	name string
	Age  int `json:"age"`
}
