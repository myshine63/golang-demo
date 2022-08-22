package main

import "fmt"

func main() {
	defer fmt.Println("exist")
	panic("123")
	defer fmt.Println("end")
}
