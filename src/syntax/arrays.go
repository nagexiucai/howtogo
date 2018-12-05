package main

import "fmt"

func main() {
	// 遍历数组
	var array = [...]int{0,1,2,3,4}
	for i:=0; i<len(array); i++ {
		fmt.Println(array[i])
	}

	// 拷贝和引用
	strArrayX := [...]string{"Xiucai", "Huaan", "Qiuxiang"}
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Println("----------")

	strArrayY := strArrayX
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Println("----------")

	strArrayX[0] = "HuaWen"
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Println("----------")

	strArrayZ := &strArrayX
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Printf("strArrayZ: %v\n", strArrayZ)
	fmt.Printf("&strArrayZ: %v\n", &strArrayZ)
	fmt.Printf("*strArrayZ: %v\n", *strArrayZ)
	fmt.Println("-----*&strArrayX-----:", *&strArrayX)
	fmt.Println("----------")

	strArrayX[1] = "HuaWu"
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Printf("strArrayZ: %v\n", strArrayZ)
	fmt.Printf("&strArrayZ: %v\n", &strArrayZ)
	fmt.Printf("*strArrayZ: %v\n", *strArrayZ)
	fmt.Println("----------")

	strArrayZ[2] = "Ningwang"
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Printf("strArrayZ: %v\n", strArrayZ)
	fmt.Printf("&strArrayZ: %v\n", &strArrayZ)
	fmt.Printf("*strArrayZ: %v\n", *strArrayZ)
	fmt.Println("----------")

	strArrayY[2] = "HuaTaiShi"
	fmt.Printf("strArrayX: %v\n", strArrayX)
	fmt.Printf("strArrayY: %v\n", strArrayY)
	fmt.Printf("strArrayZ: %v\n", strArrayZ)
	fmt.Printf("&strArrayZ: %v\n", &strArrayZ)
	fmt.Printf("*strArrayZ: %v\n", *strArrayZ)

	// 探索
	x := [...]int{9,5,2,7}
	y := [...]string{"9","5","2","7"}
	z := [...]int{0:0,4:4}

	fmt.Println(x, y, z)
}
