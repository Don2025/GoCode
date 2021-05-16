package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("year=%s\nmonth=%s\ndate=%s\n", s[:4], s[4:6], s[6:])
}
/*
 * 运行时间：6ms 超过8.00%用Go提交的代码
 * 占用内存：888KB 超过52.00%用Go提交的代码
**/