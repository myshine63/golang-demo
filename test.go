package main

import "net"

func main() {
	err := net.Dial("tcp", "8080")
}
