package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(longestCommonSubsequence(arr[0], arr[1]))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了77.63%的用户
 * 占用内存：10.5MB 在所有Go提交中击败了89.56%的用户
**/
