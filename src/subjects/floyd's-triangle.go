package main

import "fmt"

/* [Robert W. Floyd](https://www-cs-faculty.stanford.edu/~knuth/papers/floyd.ps.gz) */

func main() {
	var rows int
	var temp int = 1
	fmt.Println("Enter number of rows:")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf(" %d", temp)
			temp++
		}
		fmt.Println("")
	}
}
