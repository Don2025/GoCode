package main

import "fmt"

func main() {
	var str string
	var hash [128]bool
	var cnt int
	fmt.Scanf("%s", &str)
	for _, x := range str {
		if x >= 0 && x <= 127 {
			if !hash[x] {
				cnt++
			}
			hash[x] = true
		}
	}
	fmt.Print(cnt)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：820KB 超过97.17%用Go提交的代码
**/
