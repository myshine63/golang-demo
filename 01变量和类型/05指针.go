package main

import "fmt"

func main() {
	// go语言有自动的垃圾回收机制(也就是gc),不需要手动回收指针
	// c语言不能返回栈上指针，go语言可以。
	// 每次运行这个方法的时候，会给变量重新分配内存，因此每次打印的指针可能不一样
	// 空指针： nil
	a := 10
	var p *int // 申明指针变量,指针变量名一般以字母p开头,int 表示指针只能指向int变量的地址
	p = &a     // 给指针变量赋值，&获取变量的地址
	println(p)
	println(*p) // 获取指针指向的内容
	println(&p) // 指针也是变量，因此也能获取指针变量的地址

	// 使用new给声明的指针分配地址，并初始地址对应的值
	var ptr *int
	ptr = new(int)
	fmt.Println(*ptr)
	fmt.Println(ptr)
	*ptr = 100
	fmt.Println(*ptr)
	fmt.Println(ptr)
}
