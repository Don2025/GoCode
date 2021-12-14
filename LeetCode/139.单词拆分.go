package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		wordDict := strings.Fields(input.Text())
		println(wordBreak(s, wordDict))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了78.67%的用户
**/
