package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/2AoeFn
// ------------------------剑指 Offer II Problem 98------------------------
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// ------------------------剑指 Offer II Problem 98------------------------
/*
 * https://leetcode.cn/problems/2AoeFn
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了53.90%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input two numbers separated by space on a line:")
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(arr[0])
		n, _ := strconv.Atoi(arr[1])
		Printf("Output: %v\n", uniquePaths(m, n))
		Printf("Input two numbers separated by space on a line:")
	}
}
