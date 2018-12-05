package main

import "fmt"

func main() {
	var employee = map[string]string{}
	employee["xiucai"] = "nagexiucai.com"
	employee["yp"] = "youngpython.com"

	score := make(map[string]int)
	score["alice"] = 100
	score["bob"] = 0

	for key,value := range employee {
		fmt.Println("Key:",key,"=>","Value:",value)
	}
	for key,value := range score {
		fmt.Println("Name:",key,"=>","Score:",value)
	}
	score["bob"] = 66
	score["clark"] = 77
	for key,value := range score {
		fmt.Println("Name:",key,"=>","Score:",value)
	}
	fmt.Println("length of score is", len(score))
	delete(score, "clark")
	fmt.Println("length of score is", len(score))
	delete(score, "clark")
}
