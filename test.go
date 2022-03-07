package main

import "fmt"

type User struct {
	name string
	age  int
}

var tom = User{
	name: "tom",
	age:  12,
}

func main() {
	fmt.Println(tom)
}
