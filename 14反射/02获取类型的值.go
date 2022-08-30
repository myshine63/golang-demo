package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	valueOfa := reflect.ValueOf(a)
	if valueOfa.Kind() == reflect.Int {
		fmt.Println(int(valueOfa.Int()))
		fmt.Println(valueOfa.Interface().(int))
	}
}
