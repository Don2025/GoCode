package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	if m%5 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

/*
 * 运行时间：7ms 超过20.00%用Go提交的代码
 * 占用内存：888KB 超过60.00%用Go提交的代码
**/
