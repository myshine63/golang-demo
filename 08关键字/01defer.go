package main

import (
	"fmt"
	"os"
)

// defer后面的方法会在函数结束之前调用，常用来关闭一些资源

func main() {
	redaFile("hello.txt")
}

func redaFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	file.Read(buf)
	fmt.Println(string(buf))
}
