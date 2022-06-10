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
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：5ms 超过25.00%用Go提交的代码
 * 占用内存：932KB 超过50.00%用Go提交的代码
**/
