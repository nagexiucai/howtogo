package main

import (
	"fmt"
	"sort"
)

func main() {
	var temp int
	var buff []int

	fmt.Println("请输入一些数字，回车确认，连续两个回车结束输入。")
	for {
		n, _ := fmt.Scanln(&temp)
		if n == 0 {
			break
		}
		fmt.Printf("this is %d.\n", temp)
		buff = append(buff, temp)
	}

	// 冒泡排序（大的上浮、小的下沉）
	for i := 0; i < len(buff); i++ {
		for j := i + 1; j < len(buff); j++ {
			if buff[i] < buff[j] {
				temp = buff[i]
				buff[i] = buff[j]
				buff[j] = temp
			}
		}
	}

	fmt.Println(buff)
	fmt.Println("Is Sorted:", sort.IntsAreSorted(buff))
	sort.Ints(buff) // 默认升序
	fmt.Println(buff)
}
