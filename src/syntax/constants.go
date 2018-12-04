package main

import "fmt"

func main() {
	const huaan int = 9527
	fmt.Println(huaan)

	// huaan = 1314 # can not assign to a constant
	fmt.Println(huaan)
	// const qiuxiang int // const declaration cannot have type without expression
	// const qiuxiang // missing value in const declaration
	const qiuxiang = 1314
	fmt.Println(qiuxiang)

	//const xiaoqiang := 250 // missing value in const declaration & syntax error: unexpected := at end of statement
	//fmt.Println(xiaoqiang)
}
