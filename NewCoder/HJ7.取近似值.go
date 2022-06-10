package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Print(int(n + 0.5))
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：856KB 超过89.75%用Go提交的代码
**/
