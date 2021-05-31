package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(((a % 100) + (b % 100)) % 100)
}

/*
 * 运行时间：8ms 超过7.14%用Go提交的代码
 * 占用内存：912KB 超过28.57%用Go提交的代码
**/
