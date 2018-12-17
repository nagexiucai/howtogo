package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{9, 5, 2, 7}
	fmt.Println(number)
	if sort.IntsAreSorted(number) == false {
		sort.Ints(number)
	}
	fmt.Println(number)
	fmt.Println(sort.SearchInts(number, 7))

	text := []string{"CA", "CN", "jp", "KR", "UK", "US"}
	fmt.Println(text)
	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	fmt.Println(text)
	fmt.Println(sort.SearchStrings(text, "jp"))

	vf := []float64{9.0, 5.0, 2.0, 7.0, 1314.0}
	fmt.Println(vf)
	if sort.Float64sAreSorted(vf) == false {
		sort.Float64s(vf)
	}
	fmt.Println(vf)
	fmt.Println(sort.SearchFloat64s(vf, 1314.0))
}
