package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&n)
	a = append(a, n)
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

/*
 * 运行时间：4ms 超过25.00%用Go提交的代码
 * 占用内存：928KB 超过0.00%用Go提交的代码
**/
