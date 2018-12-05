package main

import "fmt"

type MaNong struct {
}
type RECT struct {
	width float64
	length float64
	color string
}

func main() {
	rect := RECT{95.27,13.14,"yellow"}
	fmt.Println("area of rect is", rect.width*rect.length)
}
