package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/longest-common-subsequence/
//------------------------Leetcode Problem 1143------------------------
func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1143------------------------
/*
 * https://leetcode.cn/problems/longest-common-subsequence/
 * 执行用时：4ms 在所有Go提交中击败了77.63%的用户
 * 占用内存：10.5MB 在所有Go提交中击败了89.56%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var text1, text2 string
		Sscanf(scanner.Text(), "%s %s", &text1, &text2)
		Printf("Output: %v\n", longestCommonSubsequence(text1, text2))
	}
}
