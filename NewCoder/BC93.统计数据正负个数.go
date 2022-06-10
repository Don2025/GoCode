package main

import "fmt"

func main() {
	const n = 10
	a := make([]int, n)
	var positive, negative int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > 0 {
			positive++
		} else {
			negative++
		}
	}
	fmt.Printf("positive:%d\nnegative:%d\n", positive, negative)
}

/*
 * 运行时间：8ms 超过0.00%用Go提交的代码
 * 占用内存：904KB 超过100.00%用Go提交的代码
**/
