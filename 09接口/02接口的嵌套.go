package main

import "fmt"

/*
	1. 接口类型只能嵌入其他接口类型
	2. 接口类型不能简介或者直接嵌入本身
	3. 如果B嵌入了A，那么B可以赋值给A,但是A不可以赋值给B
**/

// One 定义一个接口
type One interface {
	logOne()
}

// Two 这里嵌套了上面的接口
type Two interface {
	One
	logTwo()
}

// Log 定义一个空结构体,下面实现了接口的方法
type Log struct{}

func (Log) logOne() {
	fmt.Println("log one")
}

func (Log) logTwo() {
	fmt.Println("log one")
}

// 泛型方法
func myLog(two Two) {
	two.logOne()
	two.logTwo()
}

func main() {
	a := &Log{}
	myLog(a)
	// 因为Log也实现了One
	var one One
	one = &Log{}
	one.logOne()

}
