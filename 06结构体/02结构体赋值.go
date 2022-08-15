package main

import "fmt"

type Book struct {
	id   string
	name string
}

// 不管是结构体实例还是结构体指针，都可以同过 "." 运算符访问和修改属性
func main() {
	// 创建并赋值
	a := Book{
		id: "a", // 可以只给部分属性赋值，其他属于将采用默认值
	}
	fmt.Printf("%+v\n", a)
	// 创建实例指针,
	b := &Book{
		"b", // 可以省略属性名,但是必须一次列举出所有属性的值
		"bb",
	}
	fmt.Println("b.id", b.id)     // 指针类型，也可以直接访问属性
	fmt.Println("*b.id", (*b).id) // 指针类型，也可以直接访问属性
}
