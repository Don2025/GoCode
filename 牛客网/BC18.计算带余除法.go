package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	sum, mod := a/b, a%b
	fmt.Printf("%d %d", sum, mod)
}
/*
 * 运行时间：5ms 超过16.67%用Go提交的代码
 * 占用内存：1016KB 超过0.00%用Go提交的代码
**/