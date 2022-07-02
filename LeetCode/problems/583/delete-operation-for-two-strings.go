package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/delete-operation-for-two-strings/
//------------------------Leetcode Problem 583------------------------
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

//------------------------Leetcode Problem 583------------------------
/*
 * https://leetcode.cn/problems/delete-operation-for-two-strings/
 * 执行用时：8ms 在所有Go提交中击败了58.23%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了68.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %d\n", minDistance(arr[0], arr[1]))
	}
}
