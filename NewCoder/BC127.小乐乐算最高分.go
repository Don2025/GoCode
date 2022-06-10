package main

import "fmt"

func main() {
	var n, max, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x > max {
			max = x
		}
	}
	fmt.Print(max)
}

/*
 * 运行时间：3ms 超过25.00%用Go提交的代码
 * 占用内存：836KB 超过75.00%用Go提交的代码
**/
