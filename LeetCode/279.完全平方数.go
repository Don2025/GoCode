package main

import (
	"bufio"
	"os"
	"strconv"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		println(numSquares(n))
	}
}

/*
 * 执行用时：40ms 在所有Go提交中击败了59.73%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了59.28%的用户
**/
