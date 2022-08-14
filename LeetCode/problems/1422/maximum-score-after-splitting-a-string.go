package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
//------------------------Leetcode Problem 1422------------------------
func maxScore(s string) int {
	var l, ans int
	r := strings.Count(s, "1")
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			l++
		} else {
			r--
		}
		ans = max(ans, l+r)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1422------------------------
/*
 * https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了84.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", maxScore(scanner.Text()))
	}
}
