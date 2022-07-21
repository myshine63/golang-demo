package main

import "fmt"

func log(value interface{}, key bool) {
	if key {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
	fmt.Printf("value is %v,类型是:%T \n", value, value)
}

func switchType() {
	// 通过switch-type进行类型断言
	arr := []interface{}{123, false, "hello tom"}
	for _, value := range arr {
		switch value.(type) {
		case int:
			log(value, true)
		case bool:
			log(value, true)
		case string:
			log(value, true)
		default:
			fmt.Println("错误的值")
		}

	}
}

func okPattern() {
	var a interface{} = 1
	// 断言失败，ok为false,b为字符串类型的空值
	b, ok := a.(string)
	log(b, ok)
	// 断言失败，ok为false,c为bool类型的默认值也就是false
	c, ok := a.(bool)
	log(c, ok)
	// 如果断言成功，这ok为true,b将转换成字符串类型
	d, ok := a.(int)
	log(d, ok)
}

func main() {
	switchType()
	//okPattern()
}
