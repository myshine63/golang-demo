package main

import "fmt"

func main() {
	// 切片的长度和容量可变
	//sliceByArr()
	//newSlice()
	appendData()
}

func sliceByArr() {
	// 申明一个切片
	var newSlice []int
	fmt.Println(newSlice)
	// 通过数组生成一个切片
	arr := [...]string{"tom", "is", "a", "good", "boy"}
	slice := arr[0:] //不能超过数原数组长度
	fmt.Println("复制的切片:", slice)
	// 从指定位置生产切片，不包最后
	slice2 := arr[1:3]
	fmt.Println("1-3生成的切片", slice2)
	fmt.Println("切片容量", cap(slice2))
	fmt.Println("切片长度", len(slice2))
	// 切片和原数组共享内存
	slice2[0] = "isn't"
	fmt.Println(arr)
	fmt.Println(slice)
}

func newSlice() {
	// 申明一个切片，注意没有大小，这个时候切片未分配内存，可以用nil判断
	var slice []int
	// 判断切片是否为空
	if slice == nil {
		fmt.Println("slice是个空切片")
	}
	fmt.Println(slice) // 相当于空数组
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// 不能通过a == nil判断切片是否为空
	a := []int{}
	if a == nil {
		fmt.Println("a是空切片")
	}
	// 给申明后的切片赋值
	slice = make([]int, 4, 8) // 切片的长度为4，容量为8，并且前4个元素用0进行填充
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// 申明切片并赋值
	slice1 := []int{1, 2, 3, 4, 5} //注意没有限定长度,此时len=cap=5
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}

func appendData() {
	// 注意数组不能通过append添加元素，当cap=len时再添加元素，cap=cap*2.当cap再次等于len时cap会再次扩大
	// 当切片用0进行填充后，append会在切片末尾添加元素，这时切片的长度+1
	// append添加元素，并不会修改原来切片的值
	slice := make([]int, 1, 4)
	slice1 := append(slice, 1)
	fmt.Println(&slice[0] == &slice1[0]) // true
	fmt.Println(slice)
	fmt.Println(slice1)
}
