package main

import "fmt"

// 所谓安全模式，就是容易出错的方法，都传入一个安全方法中去执行

func safe(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉了错误", err)
		}
	}()
	f()
}
func main() {
	safe(isAErr)
	fmt.Println("hello tom")
}
func isAErr() {
	panic("this is a err")
}
