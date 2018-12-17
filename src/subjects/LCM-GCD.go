package main

import "fmt"

// LCM --- Lowest Common Multiple --- 最小公倍数
// GCD --- Greatest Common Divisor --- 最大公约数

func LCM(a, b int) int {
	n := a
	// 取大者
	if a < b {
		n = b
	}
	for {
		if n%a == 0 && n%b == 0 {
			break
		}
		n++
	}
	return n
}

func GCD(a, b int) int {
	n := a
	// 取小者
	if a > b {
		n = b
	}
	for {
		if a%n == 0 && b%n == 0 {
			break
		}
		n--
	}
	return n
}

func main() {
	fmt.Println("LCM of 9527 and 1314 is", LCM(9527, 1314))
	fmt.Println("GCD of 9527 and 1314 is", GCD(9527, 1314))
}
