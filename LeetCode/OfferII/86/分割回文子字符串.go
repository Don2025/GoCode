package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/M99OJA
// ------------------------剑指 Offer II Problem 86------------------------
func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	var strs []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), strs...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				strs = append(strs, s[i:j+1])
				dfs(j + 1)
				strs = strs[:len(strs)-1]
			}
		}
	}
	dfs(0)
	return
}

// ------------------------剑指 Offer II Problem 86------------------------
/*
 * https://leetcode.cn/problems/M99OJA
 * 执行用时：284ms 在所有Go提交中击败了11.18%的用户
 * 占用内存：22MB 在所有Go提交中击败了96.27%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", partition(scanner.Text()))
	}
}
