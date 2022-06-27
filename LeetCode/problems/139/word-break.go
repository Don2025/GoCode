package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/word-break/
//------------------------Leetcode Problem 139------------------------
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, x := range wordDict {
		wordMap[x] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//------------------------Leetcode Problem 139------------------------
/*
 * https://leetcode.cn/problems/word-break/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了78.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		wordDict := strings.Fields(scanner.Text())
		Printf("Output: %v\n", wordBreak(s, wordDict))
	}
}
