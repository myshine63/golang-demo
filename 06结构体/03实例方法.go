package main

import "fmt"

type Fruits struct {
	name  string
	price int
}

/**
结构体的实例方法
*/
// 参数为实例类型，不会修改属性的值
func (fruit Fruits) changePrice(price int) {
	fruit.price = price
}

// 参数为指针类型，会修改实例的值
func (fruit *Fruits) changeName(name string) {
	fruit.name = name
}
func main() {
	apple := Fruits{
		name:  "apple",
		price: 12,
	}
	apple.changePrice(15)
	fmt.Println(apple)
	apple.changeName("orange")
	fmt.Println(apple)
}
