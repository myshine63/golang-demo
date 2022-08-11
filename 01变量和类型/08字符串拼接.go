package main

import (
	"bytes"
	"fmt"
)

func main() {
	a, b := "hello ", "tom"
	// 简单相加
	c := a + b
	fmt.Println(c)
	// 采用字节数组
	var d bytes.Buffer
	d.WriteString(a)
	d.WriteString(b)
	fmt.Println(d.String())
	// 使用sprint
	e := fmt.Sprint(a, b)
	fmt.Println(e)
}
