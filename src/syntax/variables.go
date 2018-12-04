package main

import "fmt"

var who int = 9527

func test() {
	fmt.Println(who)
}

func main() {
	test()
	who = 1314
	test()
}
