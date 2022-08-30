package main

import (
	"fmt"
	"reflect"
)

func main() {
	getType()
}

type Person struct {
	name string
}

func getType() {
	a := 123
	b := false
	c := "123"
	d := struct {
		name string
	}{}
	e := &struct {
		name string
	}{}
	var f interface{} = "123"
	g := [...]int{123}
	h := []int{123}
	i := &a
	j := func() {}
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b).Kind())
	fmt.Println(reflect.TypeOf(c).Kind())
	fmt.Println(reflect.TypeOf(d).Kind())
	fmt.Println(reflect.TypeOf(e).Kind())
	fmt.Println(reflect.TypeOf(f).Kind())
	fmt.Println(reflect.TypeOf(g).Kind())
	fmt.Println(reflect.TypeOf(h).Kind())
	fmt.Println(reflect.TypeOf(i).Kind())
	fmt.Println(reflect.TypeOf(i).Elem())
	fmt.Println(reflect.TypeOf(j).Kind())
}
