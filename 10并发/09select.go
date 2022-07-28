package main

import (
	"fmt"
	"time"
)

// select 用来监听通道的行为
// 所有case分支都必须是io
// 当其中一条case的用例可以执行时，便会执行。

func main() {
	ch := make(chan int, 10)
	// 写入数据
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
loop:
	for {
		time.Sleep(time.Second)
		fmt.Println("监听中。。。")
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-time.After(3 * time.Second):
			break loop
		}
	}
	fmt.Println("over!")
}
