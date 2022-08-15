package main

import "fmt"

func main() {
	//f()
	f2()
}

func f() {
	arr := [...]int{0, 1, 2, 3, 4}
	slice1 := arr[:2]    // 0 1
	slice2 := slice1[1:] // 1
	fmt.Println(slice2)
	slice3 := slice1[1:5] // 可以超过slice的长度，但是不能超过s1的容量
	fmt.Println(slice3)   // 1,2,3,4
}

// 当添加元素，超过slice的容量时，将会从新生成一个切片，这个切片将不在是对原数组的引用
func f2() {
	arr := [...]int{0, 1, 2}
	slice := arr[:]
	slice1 := append(slice, 3) // slice1已经不再是原数组的引用
	fmt.Println("arr", arr)
	fmt.Println("slice", slice)
	fmt.Println("slice1", slice1)
	slice[0] = 10
	slice1[0] = 10000
	fmt.Println("arr", arr)
	fmt.Println("slice", slice)
	fmt.Println("slice1", slice1)
}
