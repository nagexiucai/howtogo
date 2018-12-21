package main

import "fmt"

func Factorial(n uint64) uint64 {
	result := uint64(1)
	for i:=uint64(1); i<=n; i++ {
		result *= uint64(i)
	}
	return result
}
func main() {
	fmt.Println("9 --- factorial =", Factorial(9))
	// 溢出
	fmt.Println("9527 --- factorial =", Factorial(9527))
	fmt.Println("1314 --- factorial =", Factorial(1314))
}
