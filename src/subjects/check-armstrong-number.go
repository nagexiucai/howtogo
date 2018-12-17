package main

import "fmt"

/* Armstrong Number
 * 阿姆斯特朗数、水仙花数
 * 一个正整数等于每位数字的立方和：153=1^3+5^3+3^3
 */

func main() {
	var number, temp, remainder int
	var result int = 0
	fmt.Println("Enter any three digit number:")
	fmt.Scan(&number)
	temp = number

	for {
		remainder = temp % 10
		result += remainder * remainder * remainder
		temp /= 10

		if temp == 0 {
			break
		}
	}

	if result == number {
		fmt.Printf("%d is an armstrong number.", number)
	} else {
		fmt.Printf("%d is not an armstrong number.", number)
	}
}
