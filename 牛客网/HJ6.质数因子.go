package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}
	if n > 1 {
		fmt.Printf("%d ", n)
	}
}

/*
 * 运行时间：3ms 超过64.47%用Go提交的代码
 * 占用内存：940KB 超过44.64%用Go提交的代码
**/
