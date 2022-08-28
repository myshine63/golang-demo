package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	go pro(ch)
	cus(ch)
	//time.Sleep(time.Second)
}
func pro(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func cus(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
