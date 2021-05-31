package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Happy new year!Good luck!\n")
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：872KB 超过100.00%用Go提交的代码
**/
