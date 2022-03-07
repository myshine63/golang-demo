package main

import "fmt"

func main() {
	// if的判断条件必须强制bool类型，不能是0,"",nil,
	const a = "hello tom"
	if len(a) < 10 {
		fmt.Println("小于10")
	} else {
		fmt.Println("大于10")
	}
	// a的作用域只限于这个if语句
	if a := "123"; len(a) <= 3 {
		fmt.Println("123的长度小于3")
	} else {
		fmt.Println("123的长度大于3")
	}
}
