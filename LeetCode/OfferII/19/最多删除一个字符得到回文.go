package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/RQku0D/
//--------------------------剑指 Offer II Problem 19--------------------------
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return checkPalindrome(s, l, r-1) || checkPalindrome(s, l+1, r)
		}
	}
	return true
}

func checkPalindrome(s string, l int, r int) bool {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//--------------------------剑指 Offer II Problem 19--------------------------
/*
 * https://leetcode.cn/problems/RQku0D/
 * 执行用时：12ms 在所有Go提交中击败了97.80%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了62.10%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", validPalindrome(scanner.Text()))
	}
}
