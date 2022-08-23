package main

import "fmt"

/*
	1.defer用来关闭一些资源，一些错误处理
	2.defer后面的方法，会在函数结束前调用
	2.多个defer，会按照栈，先定义的后执行

*
*/
func main() {
	f2()
}

func f2() {
	defer fmt.Println("hello tom")
	defer fmt.Println("hello jerry")
	fmt.Println("hello spike")
}
