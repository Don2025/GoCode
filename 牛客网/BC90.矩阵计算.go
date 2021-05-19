package main

import "fmt"

func main() {
	var n, m, sum int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] > 0 {
				sum += a[i][j]
			}
		}
	}
	fmt.Print(sum)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：972KB 超过40.00%用Go提交的代码
**/
