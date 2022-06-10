package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	arr1 := make([][]int, m)
	arr2 := make([][]int, m)
	for i := 0; i < m; i++ {
		arr1[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	for i := 0; i < m; i++ {
		arr2[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}
	var cnt float64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr1[i][j] == arr2[i][j] {
				cnt++
			}
		}
	}
	fmt.Printf("%.2f\n", cnt/float64((m*n))*100)
}

/*
 * 运行时间：8ms 超过16.67%用Go提交的代码
 * 占用内存：1016KB 超过0.00%用Go提交的代码
**/
