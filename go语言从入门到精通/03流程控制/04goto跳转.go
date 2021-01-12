package main

func main() {
	for i := 0; ; {
		if i > 3 {
			goto outer // 跳转到指定位置
		}
		println(i)
		i++
	}
outer:
}
