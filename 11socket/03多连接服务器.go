package main

import (
	"log"
	"net"
)

// 处理错误
func handleErr(successMsg, errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
	log.Println(successMsg)
}

// 处理每个连接
func handleConnect(conn net.Conn) {
	defer conn.Close()
	// for循环可以不停的读取充客服端发来的数据
	for {
		// 创建buffer用来存储数据
		buf := make([]byte, 1024)
		// read会阻塞程序执行
		count, _ := conn.Read(buf) // 读取数据的长度
		str := string(buf[:count]) // 忽略换行
		log.Printf("从客服端读取的数据:%s,长度为:%v", str, count)
		// 向客服端发数据
		count, err := conn.Write([]byte("hello " + str))
		handleErr("写入数据成功", "写入数据失败", err)
	}
}
func servers() {
	// 创建一个服务器
	server, err := net.Listen("tcp", "localhost:8080")
	handleErr("启动服务成功...", "启动服务失败...", err)
	defer server.Close()
	log.Println("等待连接...")
	// for循环可以处理多个连接
	for {
		// 等待连接，阻止程序执行
		conn, err := server.Accept()
		handleErr("连接成功...", "连接失败...", err)
		go handleConnect(conn)
	}
}
func main() {
	servers()
}
