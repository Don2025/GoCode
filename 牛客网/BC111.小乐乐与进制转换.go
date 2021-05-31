package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var ans string
	fmt.Scan(&n)
	for n != 0 {
		ans = strconv.Itoa(n%6) + ans
		n /= 6
	}
	fmt.Print(ans)
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：864KB 超过100.00%用Go提交的代码
**/
