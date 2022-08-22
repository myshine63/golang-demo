package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//readFileErr()
	createErr()
}
func readFileErr() {
	// if err != nil 是go最常用的错误处理方式
	if f, err := os.ReadFile("test.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(f))
	}
}

// go中的error实际上是一个接口,定义如下
//type error interface {
//	Error() string
//}

func createErr() {
	// 使用errors包来创建错误
	err := errors.New("this is a error")
	fmt.Println(err)
	fmt.Println(err.Error())
	// 申明一个错误，0值为nil
	var err1 error
	fmt.Println(err1)
}
