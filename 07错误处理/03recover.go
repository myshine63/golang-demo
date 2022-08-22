package main

import "fmt"

// panic 相当于try-catch中，抛出错误，recover相当于捕捉错误
func main() {
	fmt.Println("程序开始")
	defer fmt.Println("程序结束")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉了错误")
			fmt.Println("hello jerry")
		}
	}()
	createErr1()
	fmt.Println("hello tom") // 这里不会被执行
}
func createErr1() {
	defer fmt.Println("退出err1")
	panic("发生错误")
}
