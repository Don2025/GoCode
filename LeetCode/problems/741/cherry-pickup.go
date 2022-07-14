package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/cherry-pickup/
//------------------------Leetcode Problem 741------------------------
func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n<<1-1; i++ {
		for x1 := min(i, n-1); x1 >= max(i-n+1, 0); x1-- {
			for x2 := min(i, n-1); x2 >= x1; x2-- {
				y1, y2 := i-x1, i-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = math.MinInt32
					continue
				}
				ans := dp[x1][x2]
				if x1 > 0 {
					ans = max(ans, dp[x1-1][x2])
				}
				if x2 > 0 {
					ans = max(ans, dp[x1][x2-1])
				}
				if x1 > 0 && x2 > 0 {
					ans = max(ans, dp[x1-1][x2-1])
				}
				ans += grid[x1][y1]
				if x1 != x2 {
					ans += grid[x2][y2]
				}
				dp[x1][x2] = ans
			}
		}
	}
	return max(dp[n-1][n-1], 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 741------------------------
/*
 * https://leetcode.cn/problems/cherry-pickup/
 * 执行用时：8ms 在所有Go提交中击败了97.62%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了73.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			grid[i] = make([]int, n)
		}
		Printf("Output: %v\n", cherryPickup(grid))
	}
}
