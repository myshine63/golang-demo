package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		func() {
			fmt.Println("goroutine,中函数调用结束")
			//return // 结束当前函数
			//os.Exit(-1) // 退出当前线程
			runtime.Goexit() // 退出当前goroutine
		}()
		fmt.Println("goroutine结束")
	}()
	fmt.Println("主函数运行")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("主函数结束")
}
