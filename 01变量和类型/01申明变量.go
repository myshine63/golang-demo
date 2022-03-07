package main

func main() {
	// 基本数据类型: int,int8,int16,int32,int64,uint8 ...
	// float32,float64
	// 在局部范围申明的变量需要被使用,申明的全局变量可以不被使用
	// 变量不允许重复申明
	var a int // 申明单个变量
	//var a =11;
	var b, c string // 申明多个相同类型的变量
	// 申明不同类型的变量
	var (
		d int
		e string
	)
	a, b, c, d, e = 2020, "hello", "tom", 2021, "happy"
	println(a)
	println(b, c)
	println(d)
	println(e)
}
