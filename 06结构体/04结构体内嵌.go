package main

import "fmt"

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
	tom = Student{
		People: People{
			name: "tom",
			age:  13,
		},
		score: 30,
	}
	fmt.Println("tom:", tom)
	fmt.Println("name", tom.getUserName())
	fmt.Println("age:", tom.age)
	fmt.Println("age:", tom.People.age) // 避免父类覆盖子类，也可以通过这种方式访问属性
	fmt.Println("score", tom.score)
}
