package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 浅拷贝
	arr1 := arr[0:]
	arr1[0] = 111
	fmt.Println(arr)
	fmt.Println(arr1)
	// 深拷贝
	arr2 := make([]int, len(arr))
	copy(arr2, arr[0:])
	arr2[1] = 222
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(reflect.TypeOf(arr2))
}
