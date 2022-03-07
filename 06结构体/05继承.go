package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getUsername() string {
	return p.name
}

type Teacher struct {
	Person
	topic string
}

func main() {
	tom := Teacher{
		Person: Person{
			name: "tom",
			age:  12,
		},
		topic: "123",
	}
	fmt.Println(tom.name)
	fmt.Println(tom.getUsername())
}
