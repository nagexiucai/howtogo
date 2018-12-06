package main

import "fmt"

type MaNong struct {
}
type RECT struct {
	width float64
	length float64
	color string
}
type Rectangle struct {
	width float64
	length float64
	color string
	geometry struct {
		area float64
		perimeter float64
	}
}

func main() {
	rect := RECT{95.27,13.14,"yellow"}
	fmt.Println("area of rect is", rect.width*rect.length)

	// TODO: 如何用字面值初始化内部结构体字段。
	// var rectangle = Rectangle{95.27,13.14,"yellow",{0.0,0.0}} // missing type in composite literal
	// var rectangle = Rectangle{95.27,13.14,"yellow"} // too few values in struct initializer
	var rectangle = Rectangle{width:95.27,length:13.14}
	fmt.Println(rectangle)

	rectanglenew := new(Rectangle)
	rectanglenew.width = 250
	rectanglenew.length = 2
	rectanglenew.color = "red"
	fmt.Println(rectanglenew.width, rectanglenew.length, rectanglenew.color)

	rectanglenewagain := new(Rectangle)
	rectanglenewagain.width = 2
	rectanglenewagain.length = 250
	rectanglenewagain.color = "blue"
	fmt.Println(rectanglenewagain.width, rectanglenewagain.length, rectanglenewagain.color)

	// 通过结构体指针访问字段
	var prect = &Rectangle{}
	prect.width = 222
	prect.length = 222
	prect.color = "black"
	fmt.Println(prect)

	// 通过结构体实例访问字段
	var prectyet = &Rectangle{}
	(*prectyet).width = 888
	(*prectyet).length = 888
	(*prectyet).color = "white"
	fmt.Println(prectyet)
}
