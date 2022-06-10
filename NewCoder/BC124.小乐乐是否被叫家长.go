package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	if sum >= 180 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：940KB 超过55.56%用Go提交的代码
**/
