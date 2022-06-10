package main

import "fmt"

func main() {
	const n = 10
	var score [n]int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &score[i])
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", score[i])
	}
}

/*
 * 运行时间：4ms 超过100.00%用Go提交的代码
 * 占用内存：1076KB 超过0.00%用Go提交的代码
**/
