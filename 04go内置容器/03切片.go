package main

import "fmt"

func main() {
	// 切片的长度和容量可变
	appendData()
}

func sliceByArr() {
	// 通过数组形成切片和原数组共享一段内存
	// 通过下标修改切片的值时，下标不能超过切片的长度-1
	arr := [...]string{"tom", "is", "a", "good", "boy"}
	arr1 := arr[1:3] //不能超过数原数组长度
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(&arr[1] == &arr1[0])
	// 可以覆盖原数组，并且数组长度可以超过原数组，超过之后cap=len
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	arr1 = append(arr1, "jerry")
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
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
	slice1 := []int{1, 2, 3, 4, 5} //注意没有限定长度,此时len=cap=5
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	// 给申明后的切片赋值
	slice = make([]int, 4, 8) // 切片的长度为4，容量为8，并且前4个元素用0进行填充
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func appendData() {
	// 注意数组不能通过append添加元素，当cap=len时再添加元素，cap=cap*2.当cap再次等于len时cap会再次扩大
	// 当切片用0进行填充后，append会在切片末尾添加元素，这时切片的长度+1
	slice := make([]int, 1, 4)
	slice1 := append(slice, 1)
	fmt.Println(&slice[0] == &slice1[0]) // false
	fmt.Println(slice)
	fmt.Println(slice1)
}
