package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j <= i; j++ {
			if i == j || j == 0 {
				a[i][j] = 1
			} else {
				a[i][j] = a[i-1][j-1] + a[i-1][j]
			}
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：11ms 超过0.00%用Go提交的代码
 * 占用内存：888KB 超过80.00%用Go提交的代码
**/
