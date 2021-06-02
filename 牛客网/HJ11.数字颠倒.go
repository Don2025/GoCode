package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := strconv.Itoa(n)
	var a []rune
	for i := len(s) - 1; i >= 0; i-- {
		a = append(a, rune(s[i]))
	}
	fmt.Print(string(a))
}

/*
 * 运行时间：3ms 超过49.56%用Go提交的代码
 * 占用内存：932KB 超过43.39%用Go提交的代码
**/
