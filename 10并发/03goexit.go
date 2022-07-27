package main

import (
	"fmt"
	"runtime"
	"time"
)

// Goexit退出当前协程
func main() {
	go func() {
		// 创建匿名函数，并在里面结束协程
		defer fmt.Println("协程退出前，把这里执行了1")
		func() {
			fmt.Println("执行匿名函数。。。。")
			defer fmt.Println("协程退出前，把这里执行了2")
			//return // 结束当前函数
			//os.Exit(-1) // 退出当前线程
			runtime.Goexit() // 退出当前goroutine
			fmt.Println("这里就不执行了")
		}()
		fmt.Println("这里就不执行了")
	}()
	fmt.Println("主函数运行")
	time.Sleep(time.Second * 1)
	//runtime.Gosched()
	fmt.Println("主函数结束")
}
