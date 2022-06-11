package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/GzCJIP/
// ------------------------剑指 Offer II Problem 88------------------------
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 88------------------------
/*
 * https://leetcode.cn/problems/GzCJIP/
 * 执行用时：4ms 在所有Go提交中击败了83.46%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了57.89%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		cost := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minCostClimbingStairs(cost))
		Printf("Input a line of numbers separated by space:")
	}
}
