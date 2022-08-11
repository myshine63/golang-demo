package main

import "fmt"

func main() {
	// go语言没有while循环，因此可以用for去替代while循环
	i := 0
	// 其他语言中的while
	//while(i<5){
	//	fmt.Println(i)
	//	i++
	//}
	// for实现while,省略了初始语句，和赋值表达式，只保留了判断条件
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
