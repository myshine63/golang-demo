package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// 浅拷贝
	arr1 := arr[:]
	fmt.Println(&arr1[0] == &arr[0]) // true
	// 深拷贝
	arr2 := make([]int, len(arr))
	copy(arr2, arr) // copy可以将一个切片复制到另一个切片
	fmt.Println("arr2的长度", len(arr2))
	fmt.Println("arr2的容量", cap(arr2))
	fmt.Println(&arr2[0] == &arr[0]) // false
}
