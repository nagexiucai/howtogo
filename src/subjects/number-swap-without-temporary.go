package main

import "fmt"

func main() {
	var first, second int
	fmt.Println("输入两个数字：")
	fmt.Scan(&first, &second)
	fmt.Println("输入的是：", first, second)
	first = first-second
	second = first+second
	first = second-first
	fmt.Println("交换后是：", first, second)
}
