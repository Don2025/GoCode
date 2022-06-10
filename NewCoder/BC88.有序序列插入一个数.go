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
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}

/*
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：984KB 超过0.00%用Go提交的代码
**/
