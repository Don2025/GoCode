package main

import (
	"bufio"
	"os"
	"strings"
)

// https://leetcode.cn/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
// ------------------------剑指 Offer I Problem 19------------------------
func isMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	dp := make([][]bool, ns+1)
	for i := range dp {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for i := 0; i <= ns; i++ {
		for j := 1; j <= np; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[ns][np]
}

// ------------------------剑指 Offer I Problem 19------------------------
/*
 * https://leetcode.cn/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了78.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		println(isMatch(arr[0], arr[1]))
	}
}
