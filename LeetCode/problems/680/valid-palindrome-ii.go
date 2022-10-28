package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/valid-palindrome-ii/
//------------------------Leetcode Problem 680------------------------
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

//------------------------Leetcode Problem 680------------------------
/*
 * https://leetcode.cn/problems/valid-palindrome-ii/
 * 执行用时：28ms 在所有Go提交中击败了7.88%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了57.68%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", validPalindrome(scanner.Text()))
	}
}
