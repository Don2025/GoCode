package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	var max, x, y int
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] > max {
				max = a[i][j]
				x, y = i, j
			}
		}
	}
	fmt.Printf("%d %d\n", x+1, y+1)
}

/*
 * 运行时间：4ms 超过50.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
