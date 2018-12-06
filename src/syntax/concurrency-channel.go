package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var goRoutine = new(sync.WaitGroup) // goroutine组锁
var projects = make(chan string, 9) // 带缓冲通道（队列）

func Q(q chan string, number int) {
	defer goRoutine.Done() // goroutine完成

	// 刚开始从通道接收不到数据就等待
	for {
		project, state := <-q
		if state==false { // 通道清空并已关闭
			fmt.Println("Exit:", number)
			break
		}
		// 开始一条通道数据处理
		fmt.Println("Start Number:", number, "Do", project)
		sleep := rand.Int63n(60)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Time to sleep", sleep, "ms")

		// 结束一条通道数据处理
		fmt.Println("End Number:", number, "Done", project)
	}
}

func main() {
	// 延迟执行
	// 先今后出
	// 打扫战场
	defer goRoutine.Wait()
	defer close(projects)

	// 伪随机种子生成处
	rand.Seed(time.Now().Unix())

	// goroutine组锁计数器初始化
	goRoutine.Add(5)
	for i:=0;i<5;i++ {
		go Q(projects, i) // goroutine执行
	}

	// 向通道发送数据
	for i:=0;i<=9;i++ {
		projects <- fmt.Sprintf("Task %d", i)
	}
}
