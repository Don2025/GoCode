package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr1 := make([][]int, n)
	for i := 0; i < n; i++ {
		arr1[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	arr2 := make([][]int, n)
	flag := true
	for i := 0; i < n; i++ {
		arr2[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr2[i][j])
			if arr1[i][j] != arr2[i][j] {
				flag = false
			}
		}
	}
	if flag {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

/*
 * 运行时间：8ms 超过0.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
