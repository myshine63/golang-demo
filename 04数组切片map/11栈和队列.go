package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4}
	val, slice1 := pop(arr[:])
	fmt.Println(val, slice1)
}

func pop(slice []int) (int, []int) {
	if len(slice) == 0 {
		return 0, slice[:]
	}
	return slice[0], slice[1:]
}
