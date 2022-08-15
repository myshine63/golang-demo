package main

import "fmt"

func main() {
	a := 1
	b := &a
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
}
