package main

import "fmt"

type Queue []int

func (q *Queue) push(v int) {
	*q = append(*q, v)
}

func (q *Queue) pop() int {
	s := (*q)[0]
	*q = (*q)[1:]
	return s
}
func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	a := Queue{1, 2, 3}
	fmt.Printf("%p\n", &a)
	a.push(4)
	fmt.Printf("%p\n", &a)
	fmt.Println(a)
	a.push(5)
	fmt.Printf("%p\n", &a)
	fmt.Println(a)
}
