package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 12 {
		fmt.Print(2)
	} else {
		fmt.Printf("%d\n", (n/12)*4+2)
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：936KB 超过37.50%用Go提交的代码
**/
