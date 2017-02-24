package main

import "fmt"

func main() {
	var i int
	var j uint

	i = -9527
	j = 9527
	fmt.Printf("i integer %d in binary is %b.\n", i, i)
	fmt.Printf("j integer %d in binary is %b.\n", j, j)

	/* 位运算中个人比较难理解的是^Integer（按位补足）
	 ** 对于有符号数——
	 **** 正数：加一添负号
	 **** 负数：忽略负号、减一消负号
	 ** 对于无符号数：按位取反
	 **/
	i = -9527
	fmt.Printf("%d 的二进制表示是 %b，^i=%b.\n", i, i, ^i)
	fmt.Printf("%d&^%d=%d&(^%d)=%b&%b=%s&%b=%d.\n", i, i, i, i, i, ^i, i, ^i, i&(^i))
	j = 13
	fmt.Printf("j=%064b, ^j=%b.\n", j, ^j)

	var s string
	s = "i love golang so much"
	fmt.Printf("s=%s.\n", s)
}
