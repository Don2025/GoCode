package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/0i0mDW/
// ------------------------剑指 Offer II Problem 99------------------------
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 99------------------------
/*
 * https://leetcode.cn/problems/0i0mDW/
 * 执行用时：4ms 在所有Go提交中击败了98.47%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了42.86%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minPathSum(grid))
		Printf("Input the number of rows of matrix:")
	}
}
