package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", a[j][i])
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：6ms 超过100.00%用Go提交的代码
 * 占用内存：940KB 超过100.00%用Go提交的代码
**/
