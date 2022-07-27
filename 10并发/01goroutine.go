package main

import (
	"fmt"
	"time"
)

// 定义两个死循环
func task1() {
	for {
		fmt.Println("this is task 1")
		time.Sleep(time.Second * 1)
	}
}

func task2() {
	for {
		fmt.Println("this is task 2")
		time.Sleep(time.Second * 2)
	}
}
func main() {
	// 在函数面前添加go关键字，表示函数以协程的方式运行
	go task1() // 如果不添加go关键字，由于task1是死循环，则下面的代码都不会执行
	go task2()
	// 当主线程执行完毕，会终止所以协程，因此要保证主程序一直运行
	for {
		fmt.Println("this is main") // 这里最会先打印
		time.Sleep(time.Second * 3)
	}
}
