package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(f(x))
}

func f(x int) int {
	if x < 0 {
		return 1
	} else if x > 0 {
		return -1
	} else {
		return 0
	}
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：936KB 超过66.67%用Go提交的代码
**/
