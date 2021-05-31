package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if !m[a[i]] {
			fmt.Printf("%d ", a[i])
			m[a[i]] = true
		}
	}
}

/*
 * 运行时间：5ms 超过0.00%用Go提交的代码
 * 占用内存：996KB 超过50.00%用Go提交的代码
**/
