package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/stickers-to-spell-word/
//------------------------Leetcode Problem 691------------------------
func minStickers(stickers []string, target string) int {
	n := len(target)
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	var dfs func(int) int
	dfs = func(mask int) int {
		if dp[mask] != -1 {
			return dp[mask]
		}
		dp[mask] = n + 1
		for _, sticker := range stickers {
			left := mask
			cnt := make([]int, 26)
			for _, x := range sticker {
				cnt[x-'a']++
			}
			for i, x := range target {
				if mask>>i&1 == 1 && cnt[x-'a'] > 0 {
					cnt[x-'a']--
					left ^= 1 << i
				}
			}
			if left < mask {
				dp[mask] = min(dp[mask], dfs(left)+1)
			}
		}
		return dp[mask]
	}
	ans := dfs(1<<n - 1)
	if ans > n {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 691------------------------
/*
 * https://leetcode.cn/problems/stickers-to-spell-word/
 * 执行用时：96ms 在所有Go提交中击败了23.81%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了95.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stickers := strings.Fields(scanner.Text())
		scanner.Scan()
		target := scanner.Text()
		Printf("Output: %v\n", minStickers(stickers, target))
	}
}
