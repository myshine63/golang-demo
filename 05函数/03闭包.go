package main

func main() {
	// 闭包:在f1作用域定义的一个变量a，被f2作用域使用后。当f1执行完毕后，变量a的内存不会被释放，
	// 它会等到f2执行完毕后在释放，这样就形成了闭包。闭包扩展了变量的使用范围，但是资源一直得不到释放，容易引起内存泄露
	count := a(10)
	println(count(1))
	println(count(1))
}

// 在a函数定义的变量sum被f1使用了，因此当a执行完毕后，sum变量会持续存在
func a(initialVal int) func(int) int {
	f1 := func(num int) int {
		initialVal = initialVal + num
		return initialVal
	}
	return f1
}
