package main

import (
	"fmt"
	"log"
	"net"
)

func server() {
	// 创建一个监听
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("启动服务失败", err)
	}
	log.Println("启动服务成功.....")
	// 创建一个连接
	conn, err := server.Accept()
	if err != nil {
		log.Fatal("连接失败", err)
	}
	log.Println("连接成功")
	// 创建buffer用来存储数据
	buf := make([]byte, 1024)
	count, _ := conn.Read(buf) // 读取数据的长度
	str := string(buf)
	fmt.Println("从客服端读取的数据长度为：", count, str)
	// 向客服端发数据
	conn.Write([]byte("hello tom"))
	// 关闭连接
	conn.Close()
}
func main() {
	server()
}
