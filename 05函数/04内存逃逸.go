package main

import "fmt"

func main() {
	ptr := f4()
	fmt.Println(ptr)
}

// 内存逃逸，会将栈上指针放进堆上指针
func f4() *int {
	a := 123
	return &a
}
