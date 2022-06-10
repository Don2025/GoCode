package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(a[x][y])
}

/*
 * 运行时间：3ms 超过40.00%用Go提交的代码
 * 占用内存：976KB 超过20.00%用Go提交的代码
**/
