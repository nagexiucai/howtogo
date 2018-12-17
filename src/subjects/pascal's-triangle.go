package main

import "fmt"

// 帕斯卡三角——杨辉三角——贾宪三角
// 二项式定理：N次二项式展开后系数的规律

func main() {
	level := 9
	temp := 0

	for i := 0; i < level; i++ {
		for j := 0; j < level-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			fmt.Printf(" %d", temp)
		}
		fmt.Println("")
	}
}
