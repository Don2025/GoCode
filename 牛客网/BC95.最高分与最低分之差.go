package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	scores := make([]int, n)
	var max, min int = 0, 100
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
		if max < scores[i] {
			max = scores[i]
		}
		if min > scores[i] {
			min = scores[i]
		}
	}
	fmt.Print(max - min)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：888KB 超过75.00%用Go提交的代码
**/
