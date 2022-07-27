package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("len=%v,cap=%v \n", len(ch), cap(ch))
		}
	}()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-ch)
	}
}
