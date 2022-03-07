package main

func main() {
	// go语言中自增和自减单独放一行，不能用来赋值，且只有a++,a--这种形式，和其他语言不一样
	a, b := 9, 5
	println(a / b) // 整数相除会忽略掉小数部分，结果为1
	println(a % b) // 4
	c, d := 0.1, 0.2
	println(c + d) // 浮点数相加，会出现精度问题，不为0.3,
	println((c + d) == 0.3)
	a++
	println(a)
	b--
	println(b)
}
