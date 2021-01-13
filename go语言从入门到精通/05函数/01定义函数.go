package main

func main() {
	a, b := 1, 2
	sum, sub := op(a, b)
	println(sum)
	println(sub)
	println(op(a, b))
	println(defaultValue())
}

// 参数的形式: 变量名 类型。参数和返回参数列表相当于申明了一些变量，在函数体内部可以直接使用。
func op(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	// 返回值的顺序可以函数返回列表不一致，以返回的顺序为准
	return sub, sum
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
	return sum
}
