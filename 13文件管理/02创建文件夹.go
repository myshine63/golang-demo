package main

import (
	"fmt"
	"os"
)

func main() {
	createSingleDir("aa", "D:\\go-work\\src\\golang-demo\\13文件管理")
	//createMutDir("cc", "D:\\go-work\\src\\golang-demo\\13文件管理\\bb")
}

// mkdir不能创建已经存在的目录，并且只能创建一级目录
func createSingleDir(dirName, path string) {
	dirPath := fmt.Sprint(path, "\\", dirName)
	err := os.Mkdir(dirPath, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
}

// 创建多级目录，当目录已经存在时，不会做修改
func createMutDir(dirName, path string) {
	dirPath := fmt.Sprint(path, "\\", dirName)
	err := os.MkdirAll(dirPath, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
}
