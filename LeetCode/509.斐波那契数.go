package main

import (
	"bufio"
	"os"
	"strconv"
)

func fib(n int) int {
	f, g := 0, 1
	for i := 0; i < n; i++ {
		g += f
		f = g - f
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
 * 占用内存：1.9MB 在所有Go提交中击败了25.61%的用户
**/
