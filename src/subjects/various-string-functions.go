package main

import (
	"fmt"
	"strings"
)

/* Go的strings包有大量处理字符串的函数。
 * Contains ContainsAny
 * Count
 * HasPrefix
 * HasSuffix
 * Index
 * Split
 * Join
 * Repeat
 * Replace
 * ToLower
 * ToUpper
 * EqualFold -- 以UTF8编码的字符串通过Unicode小写化后比较
 * Compare
 */

func main() {
	fmt.Println("--------------------")
	fmt.Println("Contains ---", strings.Contains("China", "in"))
	fmt.Println("Contains ---", strings.Contains("Ameria", "in"))
	fmt.Println("Contains ---", strings.Contains("China", "In"))
	fmt.Println("ContainsAny ---", strings.ContainsAny("China", "In"))

	fmt.Println("--------------------")
	fmt.Println("Count ---", strings.Count("China", "in"))
	fmt.Println("Count ---", strings.Count("Ameria", "in"))

	fmt.Println("--------------------")
	fmt.Println("HasPrefix ---", strings.HasPrefix("China", "Chi"))
	fmt.Println("HasPrefix ---", strings.HasPrefix("China", "chi"))

	fmt.Println("--------------------")
	fmt.Println("HasSuffix ---", strings.HasSuffix("America", "ica"))
	fmt.Println("HasSuffix ---", strings.HasSuffix("America", "Ica"))

	fmt.Println("--------------------")
	fmt.Println("Index ---", strings.Index("China", "i"))
	fmt.Println("Index ---", strings.Index("America", "I"))

	fmt.Println("--------------------")
	fmt.Println("Split ---", strings.Split("China", "i"))
	fmt.Println("Split ---", strings.Split("China", "I"))
	fmt.Println("Split ---", strings.Split("China", ""))
	for i, item := range strings.Split("America", "") {
		fmt.Println(i, "---", item)
	}

	fmt.Println("--------------------")
	fmt.Println("Join ---", strings.Join([]string{"China", "America"}, " VS "))

	fmt.Println("--------------------")
	fmt.Println("Repeat ---", strings.Repeat("China", 9))

	fmt.Println("--------------------")
	fmt.Println("Replace ---", strings.Replace("America", "i", "I", strings.Count("America", "i")))

	fmt.Println("--------------------")
	fmt.Println("ToLower ---", strings.ToLower("America"))
	fmt.Println("ToUpper ---", strings.ToUpper("China"))

	fmt.Println("--------------------")
	fmt.Println("EqualFold ---", strings.EqualFold("Cat", "cAt"))
	fmt.Println("EqualFold ---", strings.EqualFold("Dog", "doggy"))

	fmt.Println("--------------------")
	fmt.Println("Compare ---", strings.Compare("India", "Indiana"))
	fmt.Println("Compare ---", strings.Compare("Indiana", "India"))
	fmt.Println("Compare ---", strings.Compare("India", "India"))
}
