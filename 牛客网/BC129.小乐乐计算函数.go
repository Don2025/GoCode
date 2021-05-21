package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	ans := max3(a+b, b, c) / (max3(a, b+c, c) + max3(a, b, b+c))
	fmt.Printf("%.2f\n", ans)
}

func max3(a, b, c float64) float64 {
	if a < b {
		a = b
	}
	if a < c {
		a = c
	}
	return a
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：972KB 超过0.00%用Go提交的代码
**/
