package main

import "fmt"

// interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为
// go 允许不带任何方法的 interface ，这种类型的 interface 叫 empty interface
// 所有的类型都实现了empty interface
// 如果一个类型实现了一个 interface 中所有方法，我们说类型实现了该 interface

// People 定义一个接口
type People interface {
	getName() string
	setName(name string)
}

// Person 定义一个结构体，他的两个方法和接口People的方法一样，因此他实现了People接口
type Person struct {
	name string
}

func (user Person) getName() string {
	return user.name
}

func (user Person) setName(name string) {
	user.name = name
}

// 参数为接口类型，只要实现了该接口的数据类型，都可以作为参数传进去
func changeName(people People) {
	fmt.Println("修改之前people的名字为", people.getName())
	people.setName("tom")
	fmt.Println("修改之后people的名字为", people.getName())
}

func main() {
	jerry := Person{
		name: "jerry",
	}
	// 只能将结构体实例指针赋值给接口实例
	changeName(jerry)
	// 接口类型，可以用来接收实现了该接口类型数据的指针
	var spike People // 申明一个接口实例
	spike = Person{
		name: "spike",
	}
	fmt.Println(spike.getName())
}
