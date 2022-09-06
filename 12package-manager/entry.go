package main

import (
	"fmt"
	"go.uber.org/zap"
	"golang-demo/12package-manager/calc"
	h "golang-demo/12package-manager/hello"  // 使用包别名，避免包名重复，使用别名后，只能通过别名调用包中的方法
	_ "golang-demo/12package-manager/unname" // 匿名包，不需要调用包中的方法，会自动调用包中的初始化方法
)

func main() {
	// 导入自定义包中的方法
	fmt.Println(calc.Add(1, 2, 3))
	logger, _ := zap.NewProduction()
	logger.Log(1, "导入外部的包")
	h.SayHello()
}
