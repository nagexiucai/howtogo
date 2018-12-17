package main

import "fmt"

func main() {
	fmt.Println("=====九九乘法口诀=====")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j != 1 {
				fmt.Print("\t")
			}
			fmt.Printf("%d * %d = %d", j, i, i*j)
		}
		fmt.Println()
	}
}
