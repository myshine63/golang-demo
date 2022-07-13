package main

import "fmt"

const d = 3.14 // 定义在函数外面的变量，在这个包内都可以使用
func main() {
	// 1.常量可以定义了，不被使用
	const a = "hello word"
	const (
		b = "hello tom"
		c // c的值等于上一行表达式的值，注意是表达式的值，因此可能与上一行的值不一样
	)
	fmt.Println(a, b, c)
}
