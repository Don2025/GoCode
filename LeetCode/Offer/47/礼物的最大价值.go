package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
// ------------------------剑指 Offer I Problem 47------------------------
func maxValue(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n+1)
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i-1][j-1]
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer I Problem 47------------------------
/*
 * https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
 * 执行用时：4ms 在所有Go提交中击败了97.96%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了61.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maxValue(grid))
	}
}
