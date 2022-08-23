package main

import "fmt"

// 将接口类型变量，赋值给普通类型时，需要使用类型断言

func log(value interface{}, key bool) {
	if key {
		fmt.Println("success")
		fmt.Printf("value is %v,类型是:%T \n", value, value)
	} else {
		fmt.Println("fail", value)
	}
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
	// 断言成功，返回该类型的值和true，断言失败，返回该类型的零值和false
	var a interface{} = 1
	b, ok := a.(string) // "",false
	log(b, ok)
	c, ok := a.(bool) // false,false
	log(c, ok)
	d, ok := a.(int) // 1,true
	log(d, ok)
}

func main() {
	//switchType()
	okPattern()
}
