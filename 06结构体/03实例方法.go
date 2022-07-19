package main

import "fmt"

type Fruits struct {
	name  string
	price int
}

// 指针类型实例和普通类型实例都可以通过点运算，调用实例方法

// 参数为指针类型，会改变实例的值
func (fruit *Fruits) changePrice(price int) {
	fruit.price = price
}

// 参数为非指针类型，
func (fruit Fruits) changeName(name string) {
	fruit.name = name
}
func main() {
	// 实例类型
	apple := Fruits{
		name:  "apple",
		price: 0,
	}
	// 通过实例调用方法
	apple.changeName("orange")
	apple.changePrice(15)
	fmt.Println(apple)
	// 指针类型
	pair := &Fruits{
		name:  "pair",
		price: 0,
	}
	// 通过指针类型实例调用方法
	pair.changeName("beer")
	pair.changePrice(15)
	fmt.Println(pair)
}
