package main

import "fmt"

func main() {
	student := map[string]int{
		"tom":   1,
		"jerry": 2,
		"spike": 3,
	}
	// 遍历map对象
	for key, value := range student {
		fmt.Println(key, value)
	}
	// 访问一个不存在的key，将返回value类型的默认值
	fmt.Println(student["mouse"]) // 打印数字0
	// 判断map是否有某个key
	if value, exist := student["mouse"]; exist {
		fmt.Println("mouse 存在:", value)
	} else {
		fmt.Println("mouse不存在")
	}
	// 删除map的某个属性
	delete(student, "jerry")
	fmt.Println(student)
}
