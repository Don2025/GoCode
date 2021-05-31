package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &a[i])
		if err != nil {
			return
		}
	}
	sort.Ints(a)
	for i := n - 1; i >= n-5; i-- {
		fmt.Printf("%d ", a[i])
	}
}

/*
 * 运行时间：4ms 超过22.22%用Go提交的代码
 * 占用内存：916KB 超过44.44%用Go提交的代码
**/
