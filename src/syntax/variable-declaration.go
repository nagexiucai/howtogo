package main

import "fmt"

var (name string
	age int)

func main(){
	var huaan string = "zhouxingchi"
	var qiuxiang string = "gongli"
	fmt.Println(huaan, qiuxiang)

	huawen := "huawen"
	huawu := "huawu"
	fmt.Println(huawen, huawu)

	ningwang, huataishi := "fabiao", "eng"
	fmt.Println(ningwang, huataishi)

	name = "xiucai"
	age = 32
	fmt.Println(name, age)
}
