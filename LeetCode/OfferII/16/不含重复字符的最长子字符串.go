package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/wtcaE1/
// ------------------------剑指 Offer II Problem 16------------------------
func lengthOfLongestSubstring(s string) int {
	cnt := make([]int, 256)
	var start, end, ans int
	for end < len(s) {
		if cnt[s[end]] == 0 {
			cnt[s[end]]++
			end++
		} else {
			for cnt[s[end]] != 0 {
				cnt[s[start]]--
				start++
			}
		}
		ans = max(ans, end-start)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 16------------------------
/*
 * https://leetcode.cn/problems/wtcaE1/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", lengthOfLongestSubstring(scanner.Text()))
	}
}
