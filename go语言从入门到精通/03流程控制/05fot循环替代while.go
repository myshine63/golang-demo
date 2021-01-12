package main

func main() {
	// go语言没有while循环，因此可以用for去替代while循环
	i := 0
	for {
		println(i)
		i++
		if i > 5 {
			break
		}
	}
}
