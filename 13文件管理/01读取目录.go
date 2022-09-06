package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	myReadDir()
	readAll()
}

func myReadDir() {
	// 读取文件夹下面的所有文件和文件夹,返回一个文件列表切片
	files, err := os.ReadDir("D:\\go-work\\src\\golang-demo")
	if err != nil {
		fmt.Println("读取目录出错")
	} else {
		for _, file := range files {
			if file.IsDir() {
				fmt.Printf("文件夹名字：%v\n", file.Name())
			} else {
				fmt.Printf("文件名字：%v\n", file.Name())
			}
		}
	}
}

// 递归读取该文件夹下的所有文件夹和文件，包括路径本身
func readAll() {
	err := filepath.WalkDir("D:\\go-work\\src\\golang-demo\\12package-manager", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("文件夹路径:%v,文件夹名字:%v\n", path, d.Name())
		} else {
			fmt.Printf("文件路径:%v,文件名字:%v\n", path, d.Name())
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
