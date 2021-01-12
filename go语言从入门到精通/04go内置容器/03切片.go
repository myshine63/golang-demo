package main

import "fmt"

func main() {
	// 切片的长度和容量可变
	sliceByArr()
}

func sliceByArr() {
	// 通过数组形成切片和原数组共享一段内存
	arr := [...]string{"tom", "is", "a", "good", "boy"}
	arr1 := arr[1:3]
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(&arr[1] == &arr1[0])
	fmt.Println(&arr)
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	fmt.Println(arr)
	fmt.Println(arr1)
}

func newSlice() {
	// 申明一个切片，注意没有大小
	var slice []int
	fmt.Println(slice) // 相当于空数组
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// 申明切片并赋值
	slice1 := []int{1, 2, 3, 4, 5} //注意没有限定长度
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	// 给申明后的切片赋值
	slice = make([]int, 4, 8) // 切片的长度为4，容量为8，并且用0进行填充
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func appendData() {
	slice := make([]int, 1, 1)
	fmt.Println(slice)
	b := append(slice, 2) // append会返回一个新的slice
	fmt.Println(slice)
	fmt.Println(b)
	b[0] = 123
	fmt.Println(slice[0])
}
