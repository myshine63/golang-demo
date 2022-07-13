package main

import "fmt"

func main() {
	// 特殊常量iota遇到const时，会被重置为0,没定义一个常量iota自增1
	const (
		a = 1    // 这里iota = 0
		b = iota // 这里iota = 1
		c        // c的值为上一回表达式，这里相当于c=iota
		d = 123
		e = iota
	)
	const f = iota
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
