package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/gaM7Ch/
//------------------------剑指 Offer II Problem 103------------------------
func coinChange(coins []int, amount int) int {
	maxVal := amount + 1
	dp := make([]int, maxVal)
	for i := 1; i < len(dp); i++ {
		dp[i] = maxVal
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------剑指 Offer II Problem 103------------------------
/*
 * https://leetcode.cn/problems/gaM7Ch/
 * 执行用时：8ms 在所有Go提交中击败了93.02%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了58.53%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		coins := utils.StringToInts(scanner.Text())
		scanner.Scan()
		amount, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", coinChange(coins, amount))
	}
}
