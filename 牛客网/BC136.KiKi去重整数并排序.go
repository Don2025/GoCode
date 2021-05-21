package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var hash [1010]bool
	a := []int{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if !hash[x] {
			hash[x] = true
			a = append(a, x)
		}
	}
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：948KB 超过50.00%用Go提交的代码
**/
