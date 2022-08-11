package main

import "fmt"

// iota 通常用来做枚举

func main() {
	// 特殊常量iota遇到const时，会被重置为0,每定义一个常量iota自增1
	const (
		a = 1    // 这里iota = 0
		b = iota // 这里iota = 1
		c        // c的值为上一回表达式，这里相当于c=iota，iota = 2
		d = 123  // iota = 3
		e = iota // iota=4
	)
	const f = iota // 重置iota,iota = 0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
