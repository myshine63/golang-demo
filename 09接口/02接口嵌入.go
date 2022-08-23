package main

import "fmt"

type A interface {
	speak()
}

type B interface {
	A
	run()
}

func main() {
	var a A
	var b B
	a = b
	fmt.Println(a) // nil
}
