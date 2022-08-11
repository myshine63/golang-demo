package main

import "fmt"

type S struct {
	name string
	age  int
}

func main() {
	a, b, c, d := 123, "aaa", true, 10.24
	s := S{
		"",
		14,
	}
	ptr := &a
	// 直接输出
	fmt.Printf("%v\n", s)
	// 对结构体展开
	fmt.Printf("%+v\n", s)
	// 数字
	fmt.Printf("%d\n", a)
	// 字符串
	fmt.Printf("%s\n", b)
	// 给字符串带上引号
	fmt.Printf("%q\n", b)
	// bool
	fmt.Printf("%t\n", c)
	// 浮点数
	fmt.Printf("%f\n", d)
	// 变量类型
	fmt.Printf("%T\n", b)
	// 指针类型
	fmt.Printf("%p\n", ptr)
}
