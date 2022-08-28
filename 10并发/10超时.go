package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	flag := make(chan bool)
	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(data)
			case <-time.After(time.Second * 2):
				flag <- true
			}
		}
	}()
	for i := 0; i < 3; i++ {
		ch <- i
	}
	<-flag
	fmt.Println("over")
}
