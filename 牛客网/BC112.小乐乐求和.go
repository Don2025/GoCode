package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := (1 + n) * n / 2
	fmt.Print(sum)
}

/*
 * 运行时间：4ms 超过60.00%用Go提交的代码
 * 占用内存：860KB 超过80.00%用Go提交的代码
**/
