package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	sum := (a + b - c) * d
	fmt.Print(sum)
}

/*
 * 运行时间：10ms 超过0.00%用Go提交的代码
 * 占用内存：836KB 超过100.00%用Go提交的代码
**/
