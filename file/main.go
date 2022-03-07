package main

import (
	"file/calc"
	"fmt"
)

func main() {
	// calc表示包名，Add表示包中的方法
	a := calc.Add(1, 3)
	fmt.Println(a)
}
