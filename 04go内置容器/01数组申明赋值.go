package main

import "fmt"

func main() {
	// go语言中数组只能放相同类型的元素，且长度一旦申明就固定
	var arr [5]string                                    // 长度为5，元素类型为字符串
	arr1 := [10]string{"tom", "is", "a", "good", "boy"}  // 长度为10的数组
	arr2 := [...]string{"tom", "is", "a", "good", "boy"} // 自动获取数组长度
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr = [5]string{"tom", "is", "a", "good"} // arr已经申明，因此不能用'...'
	fmt.Println(arr)
}
