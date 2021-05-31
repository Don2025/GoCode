package main

import "fmt"

func main() {
	var n float32
	fmt.Scan(&n)
	fmt.Println(int(n) % 10)
}

/*
 * 运行时间：4ms 超过25.00%用Go提交的代码
 * 占用内存：956KB 超过25.00%用Go提交的代码
**/
