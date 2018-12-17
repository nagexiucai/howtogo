package main

import (
	"fmt"
	"sort"
)

// 排序方法的稳定性：指的是排序是否影响大小一样的值原有的顺序性。
// 比如[9,5,2,2_,7]，其中“2”和“2_”都是二，为了区分这么写的。
// 排序后：[2,2_,5,7,9]，则是稳定的；
// 排序后：[2_,2,5,7,9]，则是不稳定的。
// Go的sort包中Sort方法是不稳定的，需要稳定排序，可用改包的Stable方法。

/* Sort操作的通常是Slice对象
** 需要满足三个基本接口
** Len() int
** Less(i,j int) bool
** Swap(i,j int)
** 并且可用整数索引
 */

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var ints = []int{9, 5, 2, 7}
	// what := sort.Reverse(IntSlice(ints))
	what := sort.Reverse(sort.IntSlice(ints))
	fmt.Println("After sort.Reverse(sort.IntSlice([]int))", what)
	sort.Sort(what)
	fmt.Println(what)
	fmt.Println("After sort.Sort(sort.Rverse(sort.IntSlice([]int))", ints)
	sort.Stable(IntSlice(ints))
	fmt.Println("After sort.Stable", ints)
}
