package main

import "fmt"

// 带返回值的函数。
func 加法(x int, y int) int {
	total := 0
	total = x + y
	return total
}

// 不带返回值的函数。
func 打印(a int, b int) {
	fmt.Println(a, b)
}

// 带命名返回值的函数。
func 减法(a int, b int) (z int) {
	z = a - b
	return z
}

func main() {
	fmt.Println("9527+1314 =", 加法(9527,1314))
	// fmt.Println("9527+1314 =", 加法("9527", 1314)) // cannot use "9527" (type string) as type int in argument to 加法
	打印(9527, 1314)
	fmt.Println("9527-1314 =", 减法(9527, 1314))
}
