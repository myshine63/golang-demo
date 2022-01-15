package main

import "fmt"

type Money struct {
	name string
}

type One struct {
	money Money
	price string
}

func main() {
	a := One{
		money: Money{
			name: "1块钱",
		},
		price: "1",
	}
	fmt.Print(a)
}
