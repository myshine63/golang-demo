package main

func main() {
	typeOne()
	typeTwo()
	typeThree()
	typeFour()
}

// 常规用法
func typeOne() {
	for i := 0; i < 2; i++ { // 大括号需要在同一行,另外不能加括号包起来，go会把整个括号的内容当成一个表达式
		println(i)
	}
}

func typeTwo() {
	for {
		i := 0
		for {
			if i >= 2 {
				break // 跳出内层循环
			}
			println(i)
			i++
		}
		break // 跳出外层循环
	}
}

func typeThree() {
loop:
	for {
		i := 0
		for {
			if i >= 2 {
				break loop // 结束最外层循环
			}
			println(i)
			i++
		}
	}
}

func typeFour() {
loop:
	for i := 1; i <= 3; i++ {
		for j := 0; ; j++ {
			if i == j {
				continue loop // 关闭当前循环，接着执行外层循环
			}
			println(i, j)
		}
	}
}
