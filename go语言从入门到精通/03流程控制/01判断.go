package main

import "math/rand"

func main() {
	a := rand.Intn(100)
	if !(a < 10+1) { // 大括号必须和if语句在同一行,表达式最外层没有括号
		println("success")
	} else {
		println("fail")
	}
	// 判断表达式前面可以有一个语句，其作用于为当前的if作用于内
	if a := rand.Intn(10); a < 10 {
		println("success")
	}
}
