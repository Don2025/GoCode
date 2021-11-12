package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i > 0; i-- {
		dp[i][i+1] = i
		for j := i + 1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(getMoneyAmount(n))
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了61.11%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了64.81%的用户
**/
