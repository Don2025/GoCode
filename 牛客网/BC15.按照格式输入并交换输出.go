package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("a=%d,b=%d", &a, &b)
	a, b = swap(a, b)
	fmt.Printf("a=%d,b=%d", a, b)
}

func swap(a int, b int) (int, int) {
	return b, a
}
/*
 * 运行时间：4ms 超过26.09%用Go提交的代码
 * 占用内存：956KB 超过13.04%用Go提交的代码
**/