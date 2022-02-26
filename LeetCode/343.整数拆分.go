package main

import (
	"bufio"
	"os"
	"strconv"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		var curMax int
		for j := 1; j < i; j++ {
			curMax = max(curMax, j*(i-j), j*dp[i-j])
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(a ...int) int {
	val := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > val {
			val = a[i]
		}
	}
	return val
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(integerBreak(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了93.08%的用户
**/
