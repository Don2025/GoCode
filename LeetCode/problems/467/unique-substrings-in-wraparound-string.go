package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
//------------------------Leetcode Problem 467------------------------
func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)
	var cnt, ans int
	for i := range p {
		if i > 0 && (p[i]-p[i-1]+26)%26 == 1 {
			cnt++
		} else {
			cnt = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], cnt)
	}
	for _, x := range dp {
		ans += x
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 467------------------------
/*
 * https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", findSubstringInWraproundString(scanner.Text()))
	}
}
