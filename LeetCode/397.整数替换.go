package main

import (
	"bufio"
	"os"
	"strconv"
)

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(integerReplacement(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
