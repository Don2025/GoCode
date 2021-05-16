package main

import "fmt"

func main() {
	a := [...]int{73,32,99,97,110,32,100,111,32,105,116,33}
	for _, x := range a {
		fmt.Printf(string(x))
	}
}
/*
 * 运行时间：3ms 超过0.00%用Go提交的代码
 * 占用内存：856KB 超过27.27%用Go提交的代码
**/