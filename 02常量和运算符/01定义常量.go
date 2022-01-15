package main

func main() {
	// 1.常量可以定义了，不被使用
	const a = "hello word"
	const (
		b = "hello tom"
		c // c的值等于上一行表达式的值，注意是表达式的值，因此可能与上一行的值不一样
	)
}
