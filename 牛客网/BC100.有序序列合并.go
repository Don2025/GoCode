package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, x int
	fmt.Scan(&n, &m)
	a := []int{}
	for i := 0; i < n+m; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	sort.Ints(a)
	for _, val := range a {
		fmt.Printf("%d ", val)
	}
}

/*
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：964KB 超过25.00%用Go提交的代码
**/
