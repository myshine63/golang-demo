package main

func main() {
	// 特殊常量iota遇到const时，会被重置为0,
	const (
		a = iota // 重置，因此为0
		b = "hello word"
		c        // 等于上面的表达式，因此为 hello word
		d = iota // 每定义一个常量 iota++，因此为3
	)
	const e = iota // 重置，iota为0
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
}
