package main

import (
	"bufio"
	"os"
	"strings"
)

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i, x := range word1 {
		for j, c := range word2 {
			if x == c {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return n1 + n2 - 2*dp[n1][n2]
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
		println(minDistance(arr[0], arr[1]))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了58.23%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了68.78%的用户
**/
