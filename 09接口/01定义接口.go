package main

import "fmt"

// interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为

// UserManage 定义一个UserManage接口,包含了两个方法
type UserManage interface {
	getName() string
	setName(name string)
}

// go 允许不带任何方法的 interface ，这种类型的 interface 叫 empty interface
// 所有的类型都实现了empty interface
// 如果一个类型实现了一个 interface 中所有方法，我们说类型实现了该 interface

// User 结构体实现了UserManage接口
type User struct {
	name string
}

func (user User) getName() string {
	return user.name
}

func (user *User) setName(name string) {
	user.name = name
}

// 定义一个方法，传入一个接口类型，只要实现了接口的数据类型，都可以把实例穿进去
func name(people UserManage) string {
	people.setName("jerry")
	return people.getName()
}

func main() {
	tom := User{
		"tom",
	}
	fmt.Println(name(&tom))
	fmt.Println(tom)
	var a UserManage
	a = &tom
	a.getName()
}
