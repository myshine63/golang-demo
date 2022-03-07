package main

import "fmt"

func main() {
	// 申明一个map
	var a map[string]int
	fmt.Println(a)
	//a["name"] = 123 // 没有分配空间，不能这么操作
	a = make(map[string]int, 10)
	// 增加或修改map的元素
	a["tom"] = 15
	a["jerry"] = 15
	a["spike"] = 15
	fmt.Println(a)
	fmt.Println("a的长度为", len(a))
	// 申明并赋值
	b := map[string]int{
		"tom":   15,
		"jerry": 14,
		"spike": 2020,
	}
	// 用make生成容量4的map，不能通过cap获取map的容量，可以用len(map)获取长度
	b["spike"] = 15
	fmt.Println(b)
	// 删除键值对
	delete(b, "man")
	fmt.Println(b)
	// 遍历map
	for key, value := range a {
		fmt.Println(key, value)
	}
}
