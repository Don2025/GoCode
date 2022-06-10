package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	if s1 == s2 {
		fmt.Print("same")
	} else {
		fmt.Print("different")
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：952KB 超过25.00%用Go提交的代码
**/
