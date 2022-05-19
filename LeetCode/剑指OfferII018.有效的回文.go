package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(isPalindrome(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了44.29%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了63.10%的用户
**/
