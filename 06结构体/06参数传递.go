package main

import "fmt"

type Node struct {
	id int
}

// 拷贝一个副本传递过去
func (a Node) logPoint() {
	fmt.Printf("%p\n", &a)
}

// 直接传递地址
func (a *Node) logPointByPoint() {
	fmt.Printf("%p\n", a)
}

func main() {
	root := Node{
		1,
	}
	fmt.Printf("%p\n", &root) // 和下面打印的地址一样
	root.logPointByPoint()
	root.logPoint()
}
