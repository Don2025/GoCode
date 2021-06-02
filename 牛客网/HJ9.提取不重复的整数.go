package main

import "fmt"

func main() {
	var n, sum int
	var a [11]bool
	fmt.Scanf("%d", &n)
	for i := n; i != 0; i /= 10 {
		t := i % 10
		if !a[t] {
			a[t] = true
			sum = sum*10 + t
		}
	}
	fmt.Printf("%d", sum)
}

/*
 * 运行时间：3ms 超过51.34%用Go提交的代码
 * 占用内存：900KB 超过54.02%用Go提交的代码
**/
