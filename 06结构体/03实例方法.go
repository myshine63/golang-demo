package main

import "fmt"

type Fruits struct {
	name  string
	price int
}

// 结构体实例和指针类型实例都可以通过点运算，调用实例方法

// 接受者类型为指针，会改变实例的值,将结构体的地址进行传递
func (fruit *Fruits) changePrice(price int) {
	fruit.price = price
}

// 接受者类型为值，不会改变值，将结构体拷贝一份进行传递
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
