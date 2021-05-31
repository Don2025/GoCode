package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)
	a := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		a[t] = true
	}
	for i := 0; i < n; i++ {
		if a[i] {
			fmt.Printf("%d ", i)
		}
	}
}

/*
 * 运行时间：413ms 超过25.00%用Go提交的代码
 * 占用内存：5676KB 超过75.00%用Go提交的代码
**/
