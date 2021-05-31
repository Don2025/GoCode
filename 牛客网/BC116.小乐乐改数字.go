package main

import "fmt"

func main() {
	var n, ans, x int
	x = 1
	fmt.Scan(&n)
	for n > 0 {
		t := n % 10
		if t&1 != 0 {
			ans += x
		}
		x *= 10
		n /= 10
	}
	fmt.Print(ans)
}

/*
 * 运行时间：9ms 超过0.00%用Go提交的代码
 * 占用内存：848KB 超过100.00%用Go提交的代码
**/
