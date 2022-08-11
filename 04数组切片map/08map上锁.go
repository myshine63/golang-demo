package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

func main() {
	a := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		go write(a, i)
		go read(a)
	}
	time.Sleep(time.Second * 3)
}
func write(m map[int]int, i int) {
	lock.Lock()
	m[i] = i
	lock.Unlock()
}
func read(m map[int]int) {
	lock.Lock()
	fmt.Println(m)
	lock.Unlock()
}
