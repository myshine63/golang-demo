package main

import "fmt"

// UserMessage 定义一个泛型
type UserMessage interface {
	getName() string
	setName(name string)
}

// User getName,和setName方法
type User struct {
	name string
}

func (user User) getName() string {
	return user.name
}

func (user *User) setName(name string) {
	user.name = name
}

func name(name UserMessage) string {
	name.setName("jerry")
	return name.getName()
}

func main() {
	tom := User{
		"tom",
	}
	fmt.Println(name(&tom))
	fmt.Println(tom)
}
