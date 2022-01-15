package main

import "fmt"

type Fruits struct {
	name  string
	price int
}

// 参数为指针类型
func (fruit *Fruits) changePrice(price int) {
	fruit.price = price
}

// 参数为实例类型
func (fruit Fruits) changeName(name string) {
	fruit.name = name
}
func main() {
	apple := Fruits{
		name:  "apple",
		price: 12,
	}
	apple.changePrice(111)
	fmt.Print(apple)
	apple.changeName("pia")
	fmt.Print(apple)
}
