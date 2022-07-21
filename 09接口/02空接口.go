package main

import "fmt"

// Empty 定义一个空接口
type Empty interface{}

func main() {
	var a Empty = 1
	var b Empty = "tom"
	// 因为a,b都是空接口类型，因此可以互相赋值
	a = b
	// 将空接口类型，赋值给普通类型时，需要使用类型断言
	var c int
	c = a.(int)
	fmt.Println(a, b, c)
}
