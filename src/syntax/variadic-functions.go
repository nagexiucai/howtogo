package main

import "fmt"

func main() {
	variadicExample("red", "blue", "green", "yellow")

	variadicExampleII()
	variadicExampleII("red")
	variadicExampleII("red", "blue")
	variadicExampleII("red", "blue", "green")
	variadicExampleII("red", "blue", "green", "yellow")

	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 25))
}

func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}
func variadicExampleII(s ...string) {
	fmt.Println(s)
}
func calculation(s string, y ...int) int {
	area := 1

	for _, val := range y {
		if s == "Rectangle" {
			area *= val
		} else if s == "Square" {
			area = val * val
		}
	}
	return area
}
