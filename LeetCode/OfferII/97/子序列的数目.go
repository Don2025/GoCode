package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/21dk04/
// ------------------------剑指 Offer II Problem 97------------------------
func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// ------------------------剑指 Offer II Problem 97------------------------
/*
 * https://leetcode.cn/problems/21dk04/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：10.3MB 在所有Go提交中击败了23.01%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input two strings separated by space on a line:")
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		s, t := arr[0], arr[1]
		Printf("Output: %v\n", numDistinct(s, t))
		Printf("Input two strings separated by space on a line:")
	}
}
