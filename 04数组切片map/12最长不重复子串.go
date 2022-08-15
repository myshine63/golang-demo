package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(search("abcd"))
	fmt.Println(search("abcaedfc"))
}

func search(s string) string {
	res := string(s[0])
	str := ""
	for i := 0; i < len(s); i++ {
		a := string(s[i])
		if strings.Contains(str, a) {
			if len(str) > len(res) {
				res = str
				str = a
			}
		} else {
			str += a
		}
	}
	if len(str) > len(res) {
		return str
	}
	return res
}
