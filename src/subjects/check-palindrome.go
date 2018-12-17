package main

import "fmt"

func main() {
	var inputs string
	fmt.Println("输入任意字符串：")
	fmt.Scanln(&inputs)
	fmt.Println("输入的是：", inputs)

	flag := true
	// TODO: 不支持Unicode字符
	for i := 0; i < (len(inputs) / 2); i++ {
		if inputs[i] != inputs[len(inputs)-i-1] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("是回文！")
	} else {
		fmt.Println("不是回文！")
	}
}
