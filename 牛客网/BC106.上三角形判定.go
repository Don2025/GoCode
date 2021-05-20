package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	flag := true
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
			if i > j && a[i][j] != 0 {
				flag = false
			}
		}
	}
	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
