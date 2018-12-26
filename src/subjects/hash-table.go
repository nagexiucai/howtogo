package main

import "fmt"

func main() {
	var countries map[int]string
	countries = make(map[int]string)

	countries[0] = "America"
	countries[1] = "China"
	countries[2] = "India"

	for k, v := range countries {
		fmt.Printf("Key is %d, Value is %s.\n", k, v)
	}
}
