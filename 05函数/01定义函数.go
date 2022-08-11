package main

import "fmt"

func main() {
	sum, sub := op(1, 2)
	fmt.Println(sum, sub)
	fmt.Println(defaultValue())
	fmt.Println(add(1, 2, 3))
}

// 同一个包内，函数名不能重复

// 参数的形式: 变量名 类型。参数列表和返回值列表，相当于申明了变量，可以直接在函数体内部使用
func op(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// 返回默认值
func defaultValue() (a int, b string, c bool) {
	return
}

// 可变参数。slice表示切片，用来保存数据，...int，表示可变参数的类型都是int类型
func add(slice ...int) (sum int) {
	sum = 0
	for _, value := range slice {
		sum += value
	}
	return
}
