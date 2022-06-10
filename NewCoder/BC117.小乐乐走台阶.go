package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	a[1], a[2] = 1, 2
	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Print(a[n])
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
