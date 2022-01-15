package main

import "fmt"

type User struct {
	id   string
	name string
	age  int
}

func main() {
	var a User // a是结构体实例
	fmt.Println("a", a)
	var b = new(User) // b是结构体实例指针
	fmt.Println("b", b)
	c := &User{} // c是结构体实例指针
	fmt.Println("c", c)
}
