package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < b {
		a = b
	}
	if a < c {
		a = c
	}
	fmt.Print(a)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：848KB 超过100.00%用Go提交的代码
**/
