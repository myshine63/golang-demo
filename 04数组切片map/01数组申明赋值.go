package main

import "fmt"

func main() {
	// go语言中数组只能放相同类型的元素，且长度一旦申明就固定,声明后使用默认值进行填充
	var arr [5]int
	fmt.Println("初始化的数组arr: ", arr)
	arr1 := [10]string{"tom", "is", "a", "good", "boy"} // 长度为10的数组，其余用空字符串填充
	fmt.Println(arr1)
	arr2 := [...]string{"tom", "is", "a", "good", "boy"} // 自动获取数组长度
	fmt.Println(arr2)
}
