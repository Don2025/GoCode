package main

import (
	"fmt"
)

func main() {
	const n = 10
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", a[i])
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：1096KB 超过0.00%用Go提交的代码
**/
