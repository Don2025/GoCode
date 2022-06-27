package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/coin-change/
//------------------------Leetcode Problem 322------------------------
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
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

//------------------------Leetcode Problem 322------------------------
/*
 * https://leetcode.cn/problems/coin-change/
 * 执行用时：12ms 在所有Go提交中击败了51.97%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了73.79%的用户
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
