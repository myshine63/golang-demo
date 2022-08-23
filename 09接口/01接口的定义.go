package main

import "fmt"

/*
1.接口是一组方法集合的类型
2.接口也可以不包含任何方法，这时称这个接口为空接口
3.所有的数据类型都实现了空接口，因此可以用空接口接收所有数据类型的值
4.如果一个数据类型实现了一个接口中所有方法，我们说类型实现该接口
5.接口类型的参数和变量，可以用来接收实现了该接口的数据类型的值和指针
6.如果结构体的方法的接受者都不为指针类型，则可以传入实例，否则必须传入指针，因此最好传入指针类型，避免出现错误
7.如果A接口的是B接口的子集，那么B接口类型的变量可以赋值给A接口，反之则不可以
**/

// People 定义一个接口
type People interface {
	getName() string
	setName(name string)
}

// Person 中的两个方法实现了People接口
type Person struct {
	name string
}

func (user Person) getName() string {
	return user.name
}

func (user Person) setName(name string) {
	user.name = name
}

// 接口类型参数，可以用来接收实现了该接口的数据类型的值或者指针
func changeName(people People) {
	fmt.Println("修改之前people的名字为", people.getName())
	people.setName("tom")
	fmt.Println("修改之后people的名字为", people.getName())
}

func main() {
	jerry := Person{
		name: "jerry",
	}
	changeName(&jerry)
	var spike People
	spike = &Person{
		name: "spike",
	}
	fmt.Println(spike.getName())
}
