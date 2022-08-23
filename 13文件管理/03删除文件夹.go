package main

import (
	"fmt"
	"os"
)

func main() {
	//removeEmptyDir()
	removeAllFile()
}

// 只能删除空文件夹,当文件夹存在子文件或者不存在时会报错
func removeEmptyDir() {
	err := os.Remove("D:\\go-work\\src\\golang-demo\\13文件管理\\aa")
	if err != nil {
		fmt.Println("删除文件出错", err)
	}
}

// 删除文件夹，包含其子文件
func removeAllFile() {
	err := os.RemoveAll("D:\\go-work\\src\\golang-demo\\13文件管理\\aa")
	if err != nil {
		fmt.Println("删除文件出错", err)
	}
}
