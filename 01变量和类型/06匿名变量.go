package main

func main() {
	// _ 单个下划线表示匿名变量,匿名变量不会分配内存，因此可以重复定义
	// _ 匿名变量没有分配内存，因此不能进行运算，和打印
	a, _ := 10, 11
	println(a)
	//println(_)
}
