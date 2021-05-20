package main

import "fmt"

func main() {
	var n, m, k, x, y int
	fmt.Scan(&n, &m)
	a := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Scan(&k)
	var t rune
	for i := 1; i <= k; i++ {
		fmt.Scanf("%c %d %d", &t, &x, &y)
		if t == 'r' {
			for j := 1; j <= m; j++ {
				a[x][j], a[y][j] = a[y][j], a[x][j]
			}
		} else if t == 'c' {
			for j := 1; j <= n; j++ {
				a[j][x], a[j][y] = a[j][y], a[j][x]
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：1000KB 超过0.00%用Go提交的代码
**/
