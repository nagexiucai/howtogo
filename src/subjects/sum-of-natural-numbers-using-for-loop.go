package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter a positive integer:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i

	}
	fmt.Println("Sum:", sum)
}

/* Go的for循环有三种形式：
 * for initial; condition; postpone {} 常规用法
 * for condition {} 像while一样
 * for {} 无限循环
 */
