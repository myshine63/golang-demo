package main

import "fmt"

// User 定义User结构体类型，同一个包内，不能重名
type User struct {
	Id  int // 首字母大写，表示public属性，可以在当前包外被访问
	age int
}

func main() {
	/* 实例化结构体的3种方法 **/
	// a是结构体实例
	var a User
	fmt.Println("a", a)
	// 返回User类型指针
	b := new(User)
	fmt.Println("b", b)
	fmt.Println("*b", *b)
	// 返回User类型指针，和上面方式一样
	c := &User{}
	fmt.Println("c", c)
	fmt.Println("*c", *c)

}
