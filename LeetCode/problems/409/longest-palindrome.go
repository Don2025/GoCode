package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/longest-palindrome/
//------------------------Leetcode Problem 409------------------------
func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, x := range s {
		m[x]++
	}
	var ans int
	for _, it := range m {
		ans += it / 2 * 2
	}
	if ans < len(s) {
		return ans + 1
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", longestPalindrome(scanner.Text()))
	}
}

/*
 * https://leetcode.cn/problems/longest-palindrome/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了58.67%的用户
**/
