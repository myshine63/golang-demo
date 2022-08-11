package main

import "fmt"

func main() {
	fmt.Printf("4/5=%d\n", 6/5) // 整数相除会忽略掉小数部分，结果为1
	fmt.Printf("4/5=%d\n", 4/5) // 整数相除会忽略掉小数部分，结果为0
	// go语言中自增和自减单独放一行，不能用来赋值，且只有a++,a--这种形式
	a := 10
	a--
	fmt.Printf("a--= %d\n", a)
	a++
	fmt.Printf("a++= %d\n", a)
	// 浮点数相加，会出现精度问题，不为0.3,
	f1, f2 := 0.1, 0.2
	fmt.Printf("0.1+0.2=%f\n", f1+f2)
	println(f1+f1 == 0.3)
}
