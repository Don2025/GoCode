package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/qJnOS7/
// ------------------------剑指 Offer II Problem 95------------------------
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 95------------------------
/*
 * https://leetcode.cn/problems/qJnOS7/
 * 执行用时：4ms 在所有Go提交中击败了73.91%的用户
 * 占用内存：10.5MB 在所有Go提交中击败了50.50%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input two strings separated by space on a line:")
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		text1, text2 := arr[0], arr[1]
		Printf("Output: %v\n", longestCommonSubsequence(text1, text2))
		Printf("Input two strings separated by space on a line:")
	}
}
