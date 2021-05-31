package main

import "fmt"

func main() {
	const n = 4
	a := make([]int, n)
	var max int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Print(max)
}

/*
 * 运行时间：5ms 超过0.00%用Go提交的代码
 * 占用内存：1048KB 超过0.00%用Go提交的代码
**/
