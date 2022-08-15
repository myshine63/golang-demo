package main

import "fmt"

// 通过内嵌结构体可以实现继承

type People struct {
	name string
	age  int
}

func (p People) getUserName() string {
	return p.name
}

type Student struct {
	People
	score int
}

func main() {
	var tom Student
	fmt.Println("default:", tom)
	// 实例化结构体
	tom = Student{
		People: People{
			name: "tom",
			age:  13,
		},
		score: 30,
	}
	fmt.Printf("%+v\n", tom)
	fmt.Println("name:", tom.getUserName()) // 调用方法
	fmt.Println("age:", tom.People.age)     // 避免子类覆盖父类，也可以通过这种方式访问属性
	// 实例化结构体2
	jerry := Student{}
	jerry.name = "jerry"
	jerry.age = 12
	jerry.score = 90
	fmt.Printf("%+v", jerry)
}
