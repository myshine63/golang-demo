package main

import "fmt"

func main() {
	var a map[string]int
	fmt.Printf("%p", a)
	a = make(map[string]int, 10)
	a["tom"] = 12
	a["jerry"] = 12
	b := map[string]int{
		"tom":   12,
		"jerry": 12,
	}
	b["aa"] = 12
	for s, i := range b {
		fmt.Println(s, i)
	}
	if val, ok := a["tom"]; ok {
		fmt.Println(val)
	}
	delete(a, "tom")
	if val, ok := a["tom"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("不存在tom")
	}
}
