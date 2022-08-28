package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 3)
	// 向channel中写入数据
	go func() {
		// 关闭管道，之后不可以在发送数据，但是可以读取数据
		defer close(channel)
		for i := 0; i < 3; i++ {
			channel <- i
			fmt.Println("向channel中写入数据：", i)
		}
	}()
	// 读取channel中的数据
	//go func() {
	//	// 注意只有一个参数返回
	//	for value := range channel {
	//		fmt.Println("读取channel中的数据：", value)
	//	}
	//}()
	// for-range不能知道数据已经写完，必须要主动关闭通道
	for value := range channel {
		fmt.Println("读取channel中的数据：", value)
	}
	//time.Sleep(time.Second)
}
