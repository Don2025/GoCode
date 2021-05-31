package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanf("0x%x 0%o", &a, &b)
	if err != nil {
		return
	}
	fmt.Print(a + b)
}

/*
 * 运行时间：6ms 超过0.00%用Go提交的代码
 * 占用内存：764KB 超过100.00%用Go提交的代码
**/
