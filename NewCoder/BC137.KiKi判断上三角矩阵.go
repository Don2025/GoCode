package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			if j < i && x != 0 {
				fmt.Print("NO")
				return
			}
		}
	}
	fmt.Print("YES")
}

/*
 * 运行时间：3ms 超过0.00%用Go提交的代码
 * 占用内存：872KB 超过100.00%用Go提交的代码
**/
