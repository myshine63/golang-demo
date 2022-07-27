package main

import (
	"fmt"
)

// Gosched 可以阻止当前goroutine获取cpu资源，而是将其挂起，等待执行

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("this is goroutine", i)
		}
	}()
	for i := 0; i < 2; i++ {
		// Gosched可以阻止当前代码获取cpu资源
		//runtime.Gosched() // 去过去掉，则上面的代码不能执行
		fmt.Println("this is main")
	}
}
