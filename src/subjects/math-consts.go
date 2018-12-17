package main

import (
	"fmt"
	"math"
)

/* 迷人的数学常数及其背后的故事
 * Pi：圆周率常数
 * E：自然常数
 * Phi：黄金分隔比常数
 */

func main() {
	var r float64 = 9527.0

	fmt.Printf("Area of circle with radius %f is %f.\n", r, math.Pi*r*r)
}
