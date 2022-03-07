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
	user  People
	score int
}

func main() {
	tom := Student{
		user: People{
			name: "tom",
			age:  13,
		},
		score: 30,
	}
	fmt.Println(tom)
	fmt.Println(tom.user.name)
	fmt.Println(tom.user.getUserName())
}
