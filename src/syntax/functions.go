package main

import "fmt"

func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	fmt.Println("9527+1314 =", add(9527,1314))
	// fmt.Println("9527+1314 =", add("9527", 1314)) // cannot use "9527" (type string) as type int in argument to add
}
