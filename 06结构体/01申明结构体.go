package main

import "fmt"

type User struct {
	male bool
	age  int
}

func main() {
	var a User // a是结构体实例
	fmt.Println("a", a)
	var b *User // b是实例指针，但是还没有实例化，不能直接使用
	fmt.Println("b", b)
	b = new(User) // 实例化
	fmt.Println("b", *b)
	c := &User{} // c是结构体实例指针
	fmt.Println("c", c)
	fmt.Println("c", *c)
}
