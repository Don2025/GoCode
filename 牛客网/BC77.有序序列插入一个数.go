package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var x int
	fmt.Scanf("%d", &x)
	a = append(a, x)
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

/*
 * 运行时间：5ms 超过28.57%用Go提交的代码
 * 占用内存：1076KB 超过0.00%用Go提交的代码
**/
