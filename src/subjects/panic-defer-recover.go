// 比较经典的panic-defer-recover例程

package main

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			println("-----")
			handler(err)
		}
	}()
	fun()
}

func main() { Try(func() { panic("foo") }, func(e interface{}) { println(e) }) }
