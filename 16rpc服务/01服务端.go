package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloInterface interface {
	Hello(arg string, res *string) error
}

type HelloServer struct{}

func (*HelloServer) Hello(arg string, res *string) error {
	*res = arg + " 你好！"
	return nil
}

func (receiver *HelloServer) GetRegisterName() string {
	return "Hello"
}

func (receiver *HelloServer) register() error {
	return rpc.RegisterName(receiver.GetRegisterName(), receiver)
}

func main() {
	// 注册服务
	helloServer := &HelloServer{}
	_ = helloServer.register()
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	conn, _ := server.Accept()
	rpc.ServeConn(conn)
}
