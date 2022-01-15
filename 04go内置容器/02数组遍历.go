package main

func main() {
	arr := [...]int{0, 1, 2, 3, 4}
	// 用for循环遍历
	for i := 0; i < len(arr); i++ {
		println(i, arr[i])
	}
	// 用range遍历,index表示下标，value表示对应的值
	for index, value := range arr {
		println(index, value)
	}
}
