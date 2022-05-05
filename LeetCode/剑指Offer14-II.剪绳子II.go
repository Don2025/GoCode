package main

import (
	"bufio"
	"os"
	"strconv"
)

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	const mod = 1000000007
	ans := 1
	for n > 4 {
		ans = ans * 3 % mod
		n -= 3
	}
	return ans * n % mod
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(cuttingRope(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了72.85%的用户
**/
