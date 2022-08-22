package unname

import "fmt"

// init方法会在main函数之前执行
func init() {
	fmt.Println("file a 做一些初始1")
}

func init() {
	fmt.Println("file a 做一些初始2")
}
