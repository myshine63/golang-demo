package main

import (
	"fmt"
	"os"
)

func main() {
	myCreate()
	//myRead()
	//myReadAt()
}

func myCreate() {
	// 如果文件存在则删除后新建
	file, err := os.Create("D:\\go-work\\src\\golang-demo\\13文件管理\\hello.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("创建文件失败", err)
	}
	_, err = file.WriteString("hello tom\nhello jerry\nhello spike")
	if err != nil {
		fmt.Println("写入数据失败")
	}
}
func myRead() {
	// open以 O_RDONLY方法打开文件
	file, err := os.Open("D:\\go-work\\src\\golang-demo\\13文件管理\\hello.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败")
		}
	}(file)
	if err != nil {
		fmt.Println("文件未打开", err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(string(buf))
	}
}

func myReadAt() {
	// open以 O_RDONLY方法打开文件
	file, err := os.Open("D:\\go-work\\src\\golang-demo\\13文件管理\\hello.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败")
		}
	}(file)
	if err != nil {
		fmt.Println("文件未打开", err)
	}
	buf := make([]byte, 1024)
	n, _ := file.ReadAt(buf, 5)
	fmt.Println(string(buf), n)

}
