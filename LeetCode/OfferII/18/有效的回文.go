package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/XltzEq/
// ------------------------剑指 Offer II Problem 18------------------------
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// ------------------------剑指 Offer II Problem 18------------------------
/*
 * https://leetcode.cn/problems/XltzEq/
 * 执行用时：4ms 在所有Go提交中击败了44.29%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了63.10%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", isPalindrome(scanner.Text()))
	}
}
