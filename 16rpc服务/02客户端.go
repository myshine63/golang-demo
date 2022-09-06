package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	var res string
	_ = conn.Call("Student.Hello", "tom", &res)
	fmt.Println(res)
}
