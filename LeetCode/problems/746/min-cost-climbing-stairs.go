package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/min-cost-climbing-stairs/
//------------------------Leetcode Problem 746------------------------
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 746------------------------
/*
 * https://leetcode.cn/problems/min-cost-climbing-stairs/
 * 执行用时：4ms 在所有Go提交中击败了86.55%的用户
 * 占用内存：3MB 在所有Go提交中击败了38.18%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cost := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minCostClimbingStairs(cost))
	}
}
