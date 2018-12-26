package main

import "fmt"

// 金字塔（pyramid）
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

/*
          1
         1 1
        1 2 1
       1 3 3 1
      1 4 6 4 1
     1 5 10 10 5 1
    1 6 15 20 15 6 1
   1 7 21 35 35 21 7 1
  1 8 28 56 70 56 28 8 1

Process finished with exit code 0
*/