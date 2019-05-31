package main

import (
	"fmt"
	"time"
)

var layout string

func init() {
	layout = "2016=06-26 12:05:53"
}

func nit() {
	fmt.Println("i am a nit @", time.Now().Format(layout))
}

func git() {
	fmt.Println("i am a git @", time.Now().Format(layout))
}

func test(stop <-chan struct{}) {
	fmt.Println("start @", time.Now().Format(layout))
	defer git()
	go nit()
	<- stop
	fmt.Println("end @", time.Now().Format(layout))
}

func main() {
	flag := make(chan struct{}, 1)
	flag <- struct{}{}
	test(flag)
}
