package main

func main() {
	// 不需要使用break去打断流程，当有匹配项时会自动执行完退出
	// 当前case执行完毕，想执行下一个case用例时，可以添加fallthrough
	const a = 2
	switch a {
	case 1:
		println(1)
	case 2:
		println(2)
		fallthrough
	case 3:
		println(3)
	default:
		println(4)
	}
	// 结果将打印2和3
}
