package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers []float64
	var temp, sum, mean, v, sd float64

	fmt.Println("请输入若干数字（直接回车结束数字输入）：")
	for {
		n, err := fmt.Scanln(&temp)
		if err != nil {
			break
		} else if n == 0 {
			break
		} else {
			sum += temp
			numbers = append(numbers, temp)
		}
	}

	n := len(numbers)
	if n == 0 {
		fmt.Println("没有输入任何数字！")
		return
	}

	// 平均值
	mean = sum / float64(n)
	fmt.Printf("平均值是：%f\n", mean)

	// 方差
	for i := 0; i < n; i++ {
		v += math.Pow(numbers[i]-mean, 2)
	}
	fmt.Printf("方差是：%f\n", v)

	// 标准差
	sd = math.Sqrt(v)
	fmt.Printf("标准差是：%f\n", sd)
}

/*
请输入若干数字（直接回车结束数字输入）：
9
5
2
7

平均值是：5.750000
方差是：26.750000
标准差是：5.172040
*/
