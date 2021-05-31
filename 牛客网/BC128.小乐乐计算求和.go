package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	t := 1
	for i := 1; i <= n; i++ {
		t *= i
		sum += t
	}
	fmt.Print(sum)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：1020KB 超过0.00%用Go提交的代码
**/
