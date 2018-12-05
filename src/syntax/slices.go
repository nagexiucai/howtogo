package main

import "fmt"

// append/copy 都是值传递
// slice 是引用
// 对slice做append/copy 都是把slice的值拿出来操作

func main() {
	var intSliceX = make([]int, 10, 20)
	fmt.Println("intSliceX, length is", len(intSliceX), "capacity is", cap(intSliceX))

	var intSliceY = new([20]int)[0:10]
	fmt.Println("intSliceY, length is", len(intSliceY), "capacity is", cap(intSliceY))

	var intSliceZ = []int{9527, 1314}
	var strSliceO = []string{"Xiucai", "Huaan", "Qiuxiang"}
	fmt.Println("intSliceZ, length is", len(intSliceZ), "capacity is", cap(intSliceZ))
	fmt.Println("strSliceO, length is", len(strSliceO), "capacity is", cap(strSliceO))

	intSliceP := []int{9,5,2,7}
	fmt.Println("intSliceP before appending, length is", len(intSliceP), "capacity is", cap(intSliceP))
	fmt.Println("intSliceP before appending is", intSliceP)

	intSliceP = append(intSliceP, 1,3,1,4)
	fmt.Println("intSliceP after appending, length is", len(intSliceP), "capacity is", cap(intSliceP))
	fmt.Println("intSliceP after appending is", intSliceP)

	intSliceM := []int{9,5,2,7}
	intSliceN := []int{1,3,1,4}
	copy(intSliceN, intSliceM)
	fmt.Println("intSliceM is", intSliceM, "intSliceN is", intSliceN)

	strSliceS := []string{"Bei", "Nan", "Xi", "Dong", "Zhong"}
	fmt.Println(strSliceS[0])
	fmt.Println(strSliceS[1:3])
	fmt.Println(strSliceS[0:len(strSliceS)])
	fmt.Println(strSliceS[:len(strSliceS)])
	fmt.Println(strSliceS[0:])

	intSliceI := []int{9,5,2,7}
	intSliceJ := []int{}
	intSliceK := intSliceI[1:3]
	var intSliceH = new([5]int)[0:5]
	fmt.Println("intSliceI is", intSliceI, "intSliceJ is", intSliceJ, "intSliceK is", intSliceK)
	intSliceJ = append(intSliceJ, intSliceI[:1]...)
	copy(intSliceH, intSliceI[1:])
	fmt.Println("intSliceI is", intSliceI, "intSliceJ is", intSliceJ, "intSliceK is", intSliceK)
	intSliceI[0] = 1314
	fmt.Println("intSliceI is", intSliceI, "intSliceJ is", intSliceJ, "intSliceK is", intSliceK)
	intSliceI[1] = 1314
	fmt.Println("intSliceI is", intSliceI, "intSliceJ is", intSliceJ, "intSliceK is", intSliceK)
	intSliceK[1] = 1314
	fmt.Println("intSliceI is", intSliceI, "intSliceJ is", intSliceJ, "intSliceK is", intSliceK)

	fmt.Println("intSliceH is", intSliceH, "intSliceI is", intSliceI)
	intSliceI[2] = 666
	fmt.Println("intSliceH is", intSliceH, "intSliceI is", intSliceI)
	intSliceH[0] = 999
	fmt.Println("intSliceH is", intSliceH, "intSliceI is", intSliceI)
}
