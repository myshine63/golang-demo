package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(add)
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	res := f.Call(args)
	fmt.Println(res[0].Int())
}
func add(a, b int) (c int) {
	c = a + b
	return
}
