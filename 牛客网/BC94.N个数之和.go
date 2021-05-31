package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	fmt.Print(sum)
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：940KB 超过100.00%用Go提交的代码
**/
