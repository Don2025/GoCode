package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	t := gcd(n, m)
	fmt.Printf("%d ", t+m*n/t)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：940KB 超过20.00%用Go提交的代码
**/
