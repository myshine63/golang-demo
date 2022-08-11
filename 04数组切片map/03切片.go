package main

import "fmt"

func main() {
	// 切片的长度和容量可变
	//newSlice()
	//sliceByArr()
	//appendData()
	appendSlice()
}

// 申明和赋值
func newSlice() {
	// 申明一个切片，默认值为nil，长度和容量都为0
	var slice []int
	// 判断切片是否分配内存
	if slice == nil {
		fmt.Println("slice是个空切片", slice)
		fmt.Println("slice的长度", len(slice))
		fmt.Println("slice的容量", cap(slice))
	} else {
		fmt.Println(slice) // 相当于空数组
	}
	// 给申明的切片分配内存
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

// 从数组和切片中生成新切片
func sliceByArr() {
	arr := [...]string{"tom", "is", "a", "good", "boy"}
	slice := arr[:] // 复制完整的数组，生成一个新切片
	fmt.Println("复制的切片:", slice)
	fmt.Println("复制的切片的长度:", len(slice))
	fmt.Println("复制的切片的容量:", cap(slice)) // 原数组的长度-起始位置，可以变化
	// 从指定位置生成切片，不包含后面的位置
	slice2 := arr[1:3]
	fmt.Println("1-3生成的切片", slice2)
	fmt.Println("切片容量", cap(slice2))
	fmt.Println("切片长度", len(slice2))
	// 从切片中生成切片，依然是对原数组的引用
	slice3 := slice2[1:]
	fmt.Println("从slice2生成的切片", slice3)
	fmt.Println("切片容量", cap(slice3))
	fmt.Println("切片长度", len(slice3))
	// 切片和原数组共享内存
	slice3[0] = "aaa"
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
func appendData() {
	// 1.注意数组不能通过append添加元素。
	// 2.当cap=len时，再添加元素，cap=cap*2。依次类推
	// 3.append会在切片末尾添加元素，因此要注意默认的填充值
	// 4.append添加元素，并不会修改原来切片的值，而是返回一个新的切片。新切片和旧切片，共享部分内存
	slice := make([]int, 4, 8)
	slice1 := append(slice, 1)
	fmt.Println(&slice[0] == &slice1[0]) // true
	fmt.Println(slice)
	fmt.Println(slice1)
}
func appendSlice() {
	// 如果切片是从数组或者切片中生成的，在进行元素添加时，需要考虑到对原函数或者切片的影响
	arr := [...]int{1, 2, 3}
	slice := arr[:]
	slice1 := arr[0:1]
	// 很奇怪的是，如果这里改成slice1 := append(slice, 4, 5, 6),arr不会发生改变
	slice2 := append(slice1, 4, 5)
	fmt.Println(arr)    //[1,4,5]
	fmt.Println(slice)  //[1,4,5]
	fmt.Println(slice1) //[1]
	fmt.Println(slice2) //[1,4,5]
}
