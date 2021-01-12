package main

func main() {
	// a和b的值需要时同一类型
	a := 10
	b := 12
	a, b = b, a
	println(a)
	println(b)
}
