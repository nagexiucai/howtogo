package main

import "fmt"

// 算术操作符
func arithmetic() {
	var i,j int
	var m,n float32
	var p,q complex64
	var s,t string

	i = 9527
	j = 1314
	p = 1 + 1i
	q = 1 - 1i
	m = 9725.0
	n = 1314.0
	s = "what"
	t = "'s up"

	fmt.Println("i+j = ", i+j)
	fmt.Println("i-j =", i-j)
	fmt.Println("i*j =", i*j)
	fmt.Println("i/j =", i/j)

	fmt.Println("m+n =", m+n)
	fmt.Println("m-n =", m-n)
	fmt.Println("m*n =", m*n)
	fmt.Println("m/n =", m/n)

	fmt.Println("p+q =", p+q)
	fmt.Println("p-q=", p-q)
	fmt.Println("p*q =", p*q)
	fmt.Println("p/q =", p/q)

	fmt.Println("s+t =", s+t)

	// 不同类型操作，必须显式转换成相同类型再操作。
	fmt.Println("i+m =", i+int(m))
	fmt.Println("i+p =", i+int(real(p)))
	fmt.Println("i+s =", string(i)+s)
}

// 比特操作符
// “X&^Y=Z”表示“按Y的位清零X，如果Y的位是1，则Z的对应位是0，否则Z的对应为保留成X的对应位”。
func bitwise() {
	var a uint8 = 0x5A // 01011010
	var b uint8 = 0xA5 // 10100101
	var c uint8 = 0x00 // 00000000
	var d uint8 = 0xFF // 11111111
	var e uint8 = 0x5F // 01011111
	var f uint8 = 0xF5 // 11110101
	var g uint8 = 0xAF // 10101111
	var h uint8 = 0xFA // 11111010
	var i uint16 = 0xFFFF // 11111111 11111111
	var j uint16 = 0x050A // 00000101 00001010

	fmt.Printf("a(%b)&b(%b) = %b\n", a, b, a&b)
	fmt.Printf("a(%b)|b(%b) = %b\n", a, b, a|b)
	fmt.Printf("a(%b)^b(%b) = %b\n", a, b, a^b)
	fmt.Printf("a(%b)&^b(%b) = %b\n", a, b, a&^b)

	fmt.Printf("a(%b)&c(%b) = %b\n", a, c, a&c)
	fmt.Printf("a(%b)|c(%b) = %b\n", a, c, a|c)
	fmt.Printf("a(%b)^c(%b) = %b\n", a, c, a^c)
	fmt.Printf("a(%b)&^c(%b) = %b\n", a, c, a&^c)

	fmt.Printf("a(%b)&d(%b) = %b\n", a, d, a&d)
	fmt.Printf("a(%b)|d(%b) = %b\n", a, d, a|d)
	fmt.Printf("a(%b)^d(%b) = %b\n", a, d, a^d)
	fmt.Printf("a(%b)&^d(%b) = %b\n", a, d, a&^d)

	fmt.Printf("a(%b)&e(%b) = %b\n", a, e, a&e)
	fmt.Printf("a(%b)|e(%b) = %b\n", a, e, a|e)
	fmt.Printf("a(%b)^e(%b) = %b\n", a, e, a^e)
	fmt.Printf("a(%b)&^e(%b) = %b\n", a, e, a&^e)

	fmt.Printf("a(%b)&f(%b) = %b\n", a, f, a&f)
	fmt.Printf("a(%b)|f(%b) = %b\n", a, f, a|f)
	fmt.Printf("a(%b)^f(%b) = %b\n", a, f, a^f)
	fmt.Printf("a(%b)&^f(%b) = %b\n", a, f, a&^f)

	fmt.Printf("a(%b)&g(%b) = %b\n", a, g, a&g)
	fmt.Printf("a(%b)|g(%b) = %b\n", a, g, a|g)
	fmt.Printf("a(%b)^g(%b) = %b\n", a, g, a^g)
	fmt.Printf("a(%b)&^g(%b) = %b\n", a, g, a&^g)

	fmt.Printf("a(%b)&h(%b) = %b\n", a, h, a&h)
	fmt.Printf("a(%b)|h(%b) = %b\n", a, h, a|h)
	fmt.Printf("a(%b)^h(%b) = %b\n", a, h, a^h)
	fmt.Printf("a(%b)&^h(%b) = %b\n", a, h, a&^h)

	fmt.Printf("i(%b)&j(%b) = %b\n", i, j, i&j)
	fmt.Printf("i(%b)|j(%b) = %b\n", i, j, i|j)
	fmt.Printf("i(%b)^j(%b) = %b\n", i, j, i^j)
	fmt.Printf("i(%b)&^j(%b) = %b\n", i, j, i&^j)

	// XXX: 按位操作必须作用于相同类型。
	// fmt.Printf("a(%b)&i(%b) = %b\n", a, i, a&i) // invalid operation: a & i (mismatched types uint8 and uint16)
}

// 比较操作符
// Comparison Comparision
func comparison() {
	fmt.Println(9527 > 1314)
	fmt.Println(9527.0 > 1314)
	// fmt.Println(1+1i > 0) // illegal constant expression: untyped number > untyped number
	// fmt.Println(0 > "0") // cannot convert "0" (type untyped string) to type int
	fmt.Println(0 > '0')
	fmt.Println("hi" > "bob")

	fmt.Println(0 == 0.0)
	fmt.Println(0 != 0.0)
	fmt.Println(1.0 >= 1, 1 >= 1.0, 1.0 == 1)
	fmt.Println(0.0 <= 0, 0 <= 0.0, 0.0 == 0)
}

// 逻辑操作符
// 与、或、非
func __right_operand_of_logical(r bool) bool {
	fmt.Print("俺被执行嘞！")
	return r
}
func logical() {
	var u bool = true
	var v bool = false

	// fmt.Println(9527 && 1314) // invalid operation: 9527 && 1314 (operator && not defined on untyped number)
	fmt.Println(u && v)
	fmt.Println(u || v)
	fmt.Println(!u)
	fmt.Println(!u)
	fmt.Println("=====")
	fmt.Println(u && __right_operand_of_logical(true))
	fmt.Println(u || __right_operand_of_logical(true))
	fmt.Println(v && __right_operand_of_logical(false))
	fmt.Println(v || __right_operand_of_logical(false))
	fmt.Println(!__right_operand_of_logical(true))
}

func main() {
	//arithmetic()
	//bitwise()
	//comparison()
	logical()
}