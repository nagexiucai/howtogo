package main

import "fmt"

// 处理panic的defer recover必须在panic发生之前声明

func main() {
	//defer func() {
	//	action := recover()
	//	fmt.Println("刚才恐慌的值：", action)
	//}()

	var action int
	fmt.Println("输入1表示学生、输入2表示老师。")
	fmt.Scanln(&action)
	switch action {
	case 1:
		fmt.Println("学生")
	case 2:
		fmt.Println("老师")
	default:
		panic(fmt.Sprintf("非法输入：%d", action))
	}
	fmt.Println("输入1表示美国、输入2表示英国。")
	fmt.Scanln(&action)
	switch action {
	case 1:
		fmt.Println("美国")
	case 2:
		fmt.Println("英国")
	default:
		panic(fmt.Sprintf("非法输入：%d", action))
	}

	defer func() {
		action := recover()
		fmt.Println("刚才恐慌的值：", action)
	}()
}
