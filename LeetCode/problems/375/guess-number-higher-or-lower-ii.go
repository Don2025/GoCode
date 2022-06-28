package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/guess-number-higher-or-lower-ii/
//------------------------Leetcode Problem 375------------------------
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

//------------------------Leetcode Problem 375------------------------
/*
 * https://leetcode.cn/problems/guess-number-higher-or-lower-ii/
 * 执行用时：16ms 在所有Go提交中击败了61.11%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了64.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", getMoneyAmount(n))
	}
}
