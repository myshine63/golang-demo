package main

func main() {
	// 每次运行这个方法的时候，会给变量重新分配内存，因此每次打印的指针可能不一样
	a := 10
	var p *int // 申明指针变量,指针变量名一般以字母p开头,int 表示指针只能指向int变量的地址
	p = &a     // 给指针变量赋值，&获取变量的地址
	b := 10.0
	var p1 *float64
	p1 = &b
	p2 := &b

	println(p)
	println(*p) // 获取指针指向的内容
	println(&p) // 指针也是变量，因此也能获取指针变量的地址
	println(p1)
	println(*p1)
	println(p2)
	println(*p2)
}
