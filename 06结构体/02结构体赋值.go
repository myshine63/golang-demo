package main

import "fmt"

type Book struct {
	id   string
	name string
}

func main() {
	// 创建一个结构体实例
	a := Book{
		id:   "a",
		name: "aa",
	}
	fmt.Println(a)
	// 创建一个结构体实例指针
	b := &Book{
		id:   "b",
		name: "bb",
	}
	fmt.Println(b)
	// 不管是那种类型，都可以通过 . 运算符修改值
	a.id = "1"
	b.id = "2"
	fmt.Println(a, b)
}
