package main

import "fmt"

func main() {
	student := map[string]int{
		"spike": 15,
		"tom":   12,
	}
	// value, isExit:= student["jerry"],如果存在 isExit返回true，反正返回false
	if _, isExit := student["jerry"]; isExit {
		fmt.Println("jerry存在")
	} else {
		fmt.Println("jerry不存在")
	}
}
