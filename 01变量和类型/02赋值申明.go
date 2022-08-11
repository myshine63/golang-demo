package main

// 在包内都能访问
var h = 123

func main() {
	// 申明int的变量并赋值
	var a int = 10
	// 自动获取变量类型
	var b = "hello tom"
	// 短变量申明方式，不能在全局中使用，并且一次至少要申明一个新变量
	var c rune = 'c' // 字符类型的变量
	// 一次申明多个变量
	var d, e = "hello", 123
	var f string
	f, g := "hello", "myshine63" // 至少要申明一个新变量
	println(a)
	println(b)
	println(c)
	println(d, e)
	println(f, g)
	println(h)
}
