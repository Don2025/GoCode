package main

import "fmt"

func main() {
	var n, even, odd int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println(odd, even)
}

/*
 * 运行时间：4ms 超过60.00%用Go提交的代码
 * 占用内存：848KB 超过100.00%用Go提交的代码
**/
