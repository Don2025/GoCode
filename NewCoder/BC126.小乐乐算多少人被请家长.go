package main

import "fmt"

func main() {
	const m = 3
	var n, cnt int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		t := 0
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
			t += a[i][j]
		}
		if t < 180 {
			cnt++
		}
	}
	fmt.Print(cnt)
}

/*
 * 运行时间：3ms 超过25.00%用Go提交的代码
 * 占用内存：900KB 超过50.00%用Go提交的代码
**/
