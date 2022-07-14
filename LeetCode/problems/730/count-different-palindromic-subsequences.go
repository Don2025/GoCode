package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/count-different-palindromic-subsequences/
//------------------------Leetcode Problem 730------------------------
func countPalindromicSubsequences(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([][]int, n)
	next := make([][]int, n)
	pre := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		next[i] = make([]int, 4)
		pre[i] = make([]int, 4)
	}
	pos := make([]int, 4)
	for i := 0; i < 4; i++ {
		pos[i] = -1
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 4; c++ {
			pre[i][c] = pos[c]
		}
		pos[s[i]-'a'] = i
	}
	for i := 0; i < 4; i++ {
		pos[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		for c := 0; c < 4; c++ {
			next[i][c] = pos[c]
		}
		pos[s[i]-'a'] = i
	}
	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i++ {
			if s[i] == s[j] {
				low, high := next[i][s[i]-'a'], pre[j][s[i]-'a']
				if low > high {
					dp[i][j] = (2 + dp[i+1][j-1]*2) % mod
				} else if low == high {
					dp[i][j] = (1 + dp[i+1][j-1]*2) % mod
				} else {
					dp[i][j] = (dp[i+1][j-1]*2 - dp[low+1][high-1] + mod) % mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
			}
			j++
		}
	}
	return dp[0][n-1]
}

//------------------------Leetcode Problem 953------------------------
/*
 * https://leetcode.cn/problems/count-different-palindromic-subsequences/
 * 执行用时：24ms 在所有Go提交中击败了95.12%的用户
 * 占用内存：17.4MB 在所有Go提交中击败了77.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", countPalindromicSubsequences(scanner.Text()))
	}
}
