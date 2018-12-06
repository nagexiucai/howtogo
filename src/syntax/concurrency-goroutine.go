package main

// goroutine

import (
	"sync"
	"time"
	"fmt"
)

type testConcurrency struct {
	min int
	max int
	country string
}

func printCountry(test *testConcurrency, groupTest *sync.WaitGroup) {
	for i:=test.min; i<test.max; i++ {
		time.Sleep(1*time.Millisecond)
		fmt.Println(test.country)
	}
	fmt.Println()
	groupTest.Done()
}

func main() {
	groupTest := new(sync.WaitGroup)

	// struct 实例化
	// new 出来的是指针
	// 字面值初始化的是实体
	japan := &testConcurrency{0,5,"Japan"}
	china := &testConcurrency{0,5,"China"}
	india := &testConcurrency{0,5,"India"}

	go printCountry(japan, groupTest)
	go printCountry(china, groupTest)
	go printCountry(india, groupTest)

	groupTest.Add(3)
	groupTest.Wait()
}
