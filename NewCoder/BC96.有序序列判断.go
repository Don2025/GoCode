package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var flag1, flag2 bool
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i > 0 {
			if a[i] < a[i-1] {
				flag1 = true
			} else if a[i] > a[i-1] {
				flag2 = true
			}
		}
	}
	if flag1 && flag2 {
		fmt.Print("unsorted")
	} else {
		fmt.Print("sorted")
	}
}

/*
 * 运行时间：4ms 超过0.00%用Go提交的代码
 * 占用内存：912KB 超过50.00%用Go提交的代码
**/
