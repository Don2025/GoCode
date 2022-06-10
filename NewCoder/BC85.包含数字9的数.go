package main

import "fmt"

func main() {
	var cnt int
	for i := 9; i < 2020; i++ {
		for j := i; j != 0; j /= 9 {
			if j%9 == 0 {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}

/*
 * 运行时间：3ms 超过0.00%用Go提交的代码
 * 占用内存：768KB 超过100.00%用Go提交的代码
**/
