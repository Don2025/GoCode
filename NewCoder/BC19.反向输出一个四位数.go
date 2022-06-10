package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanf("%d", &n)
	for i := 0; i < 4; i++ {
		sum = sum*10 + n%10
		n /= 10
	}
	fmt.Printf("%04d\n", sum)
}

/*
 * 运行时间：3ms 超过27.78%用Go提交的代码
 * 占用内存：1024KB 超过5.56%用Go提交的代码
**/
