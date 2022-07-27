package main

import "fmt"

// 单向通道一般用来当做函数参数
// 只能向ch中写入int型数据
func producer(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		fmt.Printf("向ch中写入数据%v \n", i)
		//fmt.Println(<-ch) 不能输出
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Printf("从ch读取数据%v \n", val)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch, 3)
	consumer(ch)
}
