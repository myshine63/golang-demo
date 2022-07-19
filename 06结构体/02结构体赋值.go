package main

import "fmt"

type Book struct {
	id   string
	name string
}

func main() {
	// 创建并赋值
	a := Book{
		id: "a", // 可以只给部分属性赋值，其他属于将采用默认值
	}
	fmt.Println("a", a)
	// 创建实例指针,
	b := &Book{
		"b", // 可以省略属性名,但是必须一次列举出所有属性的值
		"bb",
	}
	fmt.Println("b.Id", b.id)     // 指针类型，也可以直接访问属性
	fmt.Println("*b.Id", (*b).id) // 指针类型，也可以直接访问属性
}
