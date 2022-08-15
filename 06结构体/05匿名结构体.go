package main

import "fmt"

type Apple struct {
	// 内嵌匿名结构体，在使用的时候，需要再次说明类型
	Product struct {
		name string
	}
	price int
}

func main() {
	var config struct {
		name string
	}
	config.name = "config"
	fmt.Printf("apple:%+v\n", config)
	apple := Apple{
		// 使用的时候，需要再次说明匿名结构体类型
		Product: struct {
			name string
		}{
			"apple",
		},
		price: 5000,
	}
	fmt.Printf("apple:%+v\n", apple)
	// 匿名结构体，没有使用type定义
	xiaoMi := struct {
		name  string
		price int
	}{
		"小米",
		2000,
	}
	fmt.Printf("apple:%+v\n", xiaoMi)
}
