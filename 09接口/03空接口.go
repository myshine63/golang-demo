package main

import "fmt"

// Empty 定义一个空接口,空接口可以接收任何的数据类型
type Empty interface{}

func main() {
	var a Empty = 1
	var b Empty = "tom"
	// 因为a,b都是空接口类型，因此可以互相赋值
	a = b
	fmt.Printf("typeof a is %T,value is %v,\n", a, a)
	fmt.Printf("typeof b is %T,value is %v,\n", b, b)
	// 将空接口赋值给其他类型时，需要使用类型断言
	var c string = b.(string)
	fmt.Println(c)
}
