package main

import "fmt"

// panic使用场景：1.程序发送不可逆错误，继续执行会对系统造成破坏。2.发生未知的错误

func main() {
	f1()
}

// panic会中断程序执行，在其后面的代码，将不会执行
func f1() {
	defer fmt.Println("这里会在程序退出之前执行")
	panic("宕机")
	defer fmt.Println("这里不会被执行")
}
