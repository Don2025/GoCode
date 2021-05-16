package main

import "fmt"

func main() {
	var n, h, m int
	_, err := fmt.Scanf("%d %d %d", &n, &h, &m)
	if err != nil {
		return
	}
	n -= m / h
	if m%h != 0 {
		n--
	}
	fmt.Println(n)
}

/*
 * 运行时间：4ms 超过15.38%用Go提交的代码
 * 占用内存：888KB 超过38.46%用Go提交的代码
**/
