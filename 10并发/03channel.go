package main

import (
	"fmt"
	"time"
)

// channel用来实现在goroutine中通信
// 读和写的次数需要相同
// 未分配空间的chan的值为nil，不能够进行读和写，

// 向chan中写入数据
func sendData(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i // 向channel中写入数据
		fmt.Println("向channel中写入数据：", i)
	}
}

// 接受chan中的数据
func receiveData(channel chan int) {
	var data int
	for i := 0; i < 10; i++ {
		data = <-channel // 将管道中的数据写给data
		fmt.Println("从channel中读取数据：", data)
	}
}

func main() {
	// 创建一个无缓冲的管道
	//channel := make(chan int)
	// 创建有缓冲的管道
	channel := make(chan int, 10)
	go sendData(channel)
	go receiveData(channel)
	time.Sleep(3 * time.Second)
}
