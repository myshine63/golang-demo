package main

func main() {
	// 匿名函数的使用方式:
	// 1.在定时后立马使用，相当于立即执行函数。
	func() {
		println("hello tom")
	}()
	// 2.将匿名函数赋值给某个变量
	f1 := func() {
		println("hello jerry")
	}
	f1()
}
