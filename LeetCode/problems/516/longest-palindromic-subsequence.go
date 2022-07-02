package main

import (
	"bufio"
	"os"
)

// https://leetcode.cn/problems/longest-palindromic-subsequence/
//------------------------Leetcode Problem 516------------------------
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 516------------------------
/*
 * https://leetcode.cn/problems/longest-palindromic-subsequence/
 * 执行用时：20ms 在所有Go提交中击败了91.86%的用户
 * 占用内存：17MB 在所有Go提交中击败了62.05%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		println(longestPalindromeSubseq(scanner.Text()))
	}
}
