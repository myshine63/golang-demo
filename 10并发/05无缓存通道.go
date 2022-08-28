package main

import (
	"fmt"
	"time"
)

// channel用来实现在goroutine间通信
// 未分配空间的chan的值为nil，不能够进行读和写，
// 无缓冲通道的发送和接受数据是同步的

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			fmt.Printf("len=%v,cap=%v\n", len(ch), cap(ch)) // channel的长度和容量
			ch <- i                                         // 向通道中写入数据
		}
	}()
	for {
		time.Sleep(time.Second)
		data, ok := <-ch
		if ok {
			fmt.Printf("从通道中读取数据:%v,通道是否关闭:%t\n", data, !ok) // 从channel中读取数据
		} else {
			fmt.Println("closed")
			break
		}
	}
}
