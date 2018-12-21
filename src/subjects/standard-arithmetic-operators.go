package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("整数和：", 9527+1314)
	fmt.Println("整数差：", 9527-1314)
	fmt.Println("整数积：", 9527*1314)
	fmt.Println("整数商：", 9527/1314)
	fmt.Println("整数模：", 9527%1314)
	fmt.Println("整数幂：", math.Pow(9527,2), math.Pow(1314, 2))

	fmt.Println("浮点数和：", 9527.2+1314)
	fmt.Println("浮点数差：", 9527-1314.2)
	fmt.Println("浮点数积：", 9527*1314.2)
	fmt.Println("浮点数商：", 9527.2/1314)
	fmt.Println("浮点数没定义模操作！")
	fmt.Println("浮点数幂：", math.Pow(9527.2,2), math.Pow(1314, 2.2))

	fmt.Println("字符串和：", "9527"+"1314")
}
