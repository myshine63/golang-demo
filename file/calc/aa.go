package calc

import "fmt"

// 一个包可以含有多个init函数，且调用没有顺序
// init函数会在引用包的时候自动调用
// 不可以手动调用init函数
func init() {
	fmt.Println("你调用了calc包")
}

// Add 需要在其它包中使用的方法，需要大写首字母
func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
