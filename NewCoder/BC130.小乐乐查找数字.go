package main

import "fmt"

func main() {
	const maxn = 101
	var n, x int
	fmt.Scan(&n)
	a := make([]int, maxn)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		a[x]++
	}
	fmt.Scan(&x)
	fmt.Print(a[x])
}

/*
 * 运行时间：3ms 超过50.00%用Go提交的代码
 * 占用内存：888KB 超过62.50%用Go提交的代码
**/
