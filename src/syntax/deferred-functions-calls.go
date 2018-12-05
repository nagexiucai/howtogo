package main

import "fmt"

func first() {
	fmt.Print("Frist")

	for i:=5; i<10; i++ {
		defer fmt.Printf("%d", i)
	}
}

func second() {
	fmt.Print("Second")
}

func main() {
	defer second()
	first()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

/* Frist43210Second
 * defer 是栈式执行顺序（LIFO）
 * defer 是紧随声明所在函数调用结束后立即执行的
 */
