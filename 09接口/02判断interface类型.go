package main

import "fmt"

func main() {
	var a interface{} // 定义一个空类型的接口变量
	a = 12            // 给泛型赋予具体的值
	// value返回a的值，ok返回bool，判断a是否为int
	if value, ok := a.(int); ok { //判断a是否为int类型，是就打印值
		fmt.Println("a的类型为int类型，值为：", value)
	} else {
		fmt.Println("a的类型不为int")
	}
	arr := make([]interface{}, 3)
	arr[0] = 1
	arr[1] = "hello tom"
	arr[2] = true
	for _, value := range arr {
		switch value.(type) {
		case int:
			fmt.Println(value, "是int类型")
		case string:
			fmt.Println(value, "是string类型")
		case bool:
			fmt.Println(value, "是bool类型")
		default:
			fmt.Println(value, "不是指定类型")
		}
	}
}
