package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/longest-palindromic-substring/
//------------------------Leetcode Problem 5------------------------
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	var expandAroundCenter func(string, int, int) (int, int)
	expandAroundCenter = func(s string, left int, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

//------------------------Leetcode Problem 5------------------------
/*
 * https://leetcode.cn/problems/longest-palindromic-substring/
 * 执行用时：4ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了88.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", longestPalindrome(scanner.Text()))
	}
}
