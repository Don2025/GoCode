package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("%s", reverse(s))
}

func reverse(s string) string {
	var a []rune
	for i := len(s) - 1; i >= 0; i-- {
		a = append(a, rune(s[i]))
	}
	return string(a)
}

/*
 * 运行时间：4ms 超过35.71%用Go提交的代码
 * 占用内存：956KB 超过46.26%用Go提交的代码
**/
