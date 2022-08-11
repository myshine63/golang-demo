package main

import "fmt"

func main() {
	fmt.Println(log())
}

func log() string {
	// defer后面的语句，会在函数执行完毕，或者return之前执行
	// 多个defer语句，会按照栈的方式，先定义的最后执行
	defer fmt.Println("1111")
	defer fmt.Println("2222")
	fmt.Println("hello jerry")
	return "hello tom"
}
