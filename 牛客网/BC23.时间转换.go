package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)
	hours := seconds/3600
	minutes := (seconds%3600)/60
	fmt.Printf("%d %d %d", hours, minutes, seconds%60)
}
/*
 * 运行时间：4ms 超过21.43%用Go提交的代码
 * 占用内存：868KB 超过57.14%用Go提交的代码
**/