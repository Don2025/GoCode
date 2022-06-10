package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)
	for _, val := range a {
		if val != x {
			fmt.Printf("%d ", val)
		}
	}
}

/*
 * 运行时间：5ms 超过14.29%用Go提交的代码
 * 占用内存：912KB 超过57.14%用Go提交的代码
**/
