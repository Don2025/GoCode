package main

import (
	"bufio"
	"os"
	"strconv"
)

func fib(n int) int {
	const mod = 1000000007
	f, g := 0, 1
	for ; n > 0; n-- {
		/* 本质上就是f, g = g, f+g
		 * g += f
		 * f = g - f
		 */
		//这题要求取模1e9+7 (1000000007)
		f, g = g, (f+g)%mod
	}
	return f
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(fib(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了71.58%的用户
**/
