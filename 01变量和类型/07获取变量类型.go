package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b, c, d, e := 2021, "hello tom", false, 'd', 0.0
	var f = &a
	var g = &d
	fmt.Println(reflect.TypeOf(a)) // int
	fmt.Println(reflect.TypeOf(b)) // string
	fmt.Println(reflect.TypeOf(c)) // bool
	fmt.Println(reflect.TypeOf(d)) // int32
	fmt.Println(reflect.TypeOf(e)) // float64
	fmt.Println(reflect.TypeOf(f)) // *int
	fmt.Println(reflect.TypeOf(g)) // *float64
}
