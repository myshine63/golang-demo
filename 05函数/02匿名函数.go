package main

import "fmt"

func main() {
	// 匿名函数，立即执行函数
	func(str string) {
		fmt.Println("hello ", str)
	}("tom")
	// 将匿名函数赋值给变量
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(f(1, 2))
}
