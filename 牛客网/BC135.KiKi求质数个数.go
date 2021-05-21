package main

import "fmt"

func main() {
	cnt := 0
	for i := 100; i < 1000; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	fmt.Print(cnt)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
 * 运行时间：3ms 超过33.33%用Go提交的代码
 * 占用内存：876KB 超过20.00%用Go提交的代码
**/
