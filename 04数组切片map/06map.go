package main

import "fmt"

func main() {
	// 申明一个map,并没有分配内存
	var a map[string]int
	//a["name"] = 123 // 没有分配空间,不能直接使用
	a = make(map[string]int, 10) // 可以不指定容量，但是可以优化
	// 增加或修改map的元素
	a["tom"] = 1
	a["jerry"] = 2
	a["spike"] = 3
	fmt.Println("a的长度为", len(a)) // 只能获取map的长度，不能获取容量
	// 申明并赋值
	b := map[string]int{
		"tom":   15,
		"jerry": 14,
		"spike": 2020,
	}
	fmt.Println(b)
}
