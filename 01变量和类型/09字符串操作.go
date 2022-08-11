package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello tom"
	// 是否以某个字符开头
	fmt.Println(strings.HasPrefix(s, "he"))
	// 是否包含某个子串
	fmt.Println(strings.Contains(s, "tom"))
	// 查找指定字符的位置
	fmt.Println(strings.IndexRune(s, 'o'))
	// 查找指定字符串的位置
	fmt.Println(strings.Index(s, "tom"))
	// 截取字符串
	fmt.Println(s[6:])
}
