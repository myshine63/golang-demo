package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	//slice = removeFromHead(5, slice)
	//fmt.Println(slice)
	slice = removeFromHead(1, slice)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

// 删除头部的前n个元素
func removeFromHead(n int, slice []int) []int {
	if n >= len(slice) || n < 0 {
		return slice[:]
	}
	return slice[n:]
}

// 删除指定位置的几个元素
func removeByIndex(index int, n int, slice []int) []int {
	if index >= len(slice) {
		return slice
	}
	if index+n > len(slice)-1 {
		return slice[0:index]
	}
	return append(slice[0:index], slice[index+n:]...)
}

// 删除尾部的几个元素

func removeFromTail(n int, slice []int) []int {
	if n > len(slice) {
		return slice[0:0]
	}
	return slice[:len(slice)-n]
}
