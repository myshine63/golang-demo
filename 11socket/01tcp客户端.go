package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn, i int) {
	for n := 0; n < 3; n++ {
		count, _ := conn.Write([]byte("this is client" + fmt.Sprintf("%d", i)))
		fmt.Println("向服务器发送数据长度", count)
		buf := make([]byte, 1024)
		conn.Read(buf)
		fmt.Println("从服务端读取的数据", string(buf))
	}

}
func connectSuccess(i int) {
	// 每次使用dial，就相当于创建了一个连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("连接失败", err) // 会阻止程序继续运行
	}
	fmt.Println("连接成功")
	// 断开链接
	handle(conn, i)
}

func connectFail() {
	// 3秒未连接成功，则超时失败
	conn, err := net.DialTimeout("tcp", "www.baidu.com:81", 3*time.Second)
	if err != nil {
		log.Fatal("连接失败", err) // 会阻止程序继续运行
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭接口失败")
		} else {
			fmt.Println("关闭连接成功")
		}
	}(conn)
	log.Println("连接成功")
}

func main() {
	//connectFail()
	for i := 0; i < 3; i++ {
		connectSuccess(i)
	}
}
