package main

import (
	"fmt"
)

// 切片采用的是引用传递，数组是值传递

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	changeArr(arr)
	fmt.Println(arr) // 1,2,3,4,4
	slice := []int{1, 2, 3, 4, 5}
	changeSlice(slice)
	fmt.Println(slice) // 100,2,3,4,5
	// 修改数组的值
	changeSlice(arr[:])
	fmt.Println(arr) // 100,2,3,4,5
}

func changeArr(arr [5]int) {
	arr[0] = 100
}

func changeSlice(slice []int) {
	slice[0] = 100
}
