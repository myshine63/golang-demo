package calc

// 大写首字母的方法，可以在包外面进行访问

func Add(nums ...int) int {
	s := 0
	for _, val := range nums {
		s += val
	}
	return s
}
