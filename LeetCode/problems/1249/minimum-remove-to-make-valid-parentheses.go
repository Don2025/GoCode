package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/minimum-remove-to-make-valid-parentheses/
//------------------------Leetcode Problem 1249------------------------
func minRemoveToMakeValid(s string) string {
	var cnt int
	ans := []rune(s)
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
			if cnt < 0 {
				ans[i] = '0'
				cnt = 0
			}
		}
	}
	cnt = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			cnt++
		} else if s[i] == '(' {
			cnt--
			if cnt < 0 {
				ans[i] = '0'
				cnt = 0
			}
		}
	}
	return strings.ReplaceAll(string(ans), "0", "")
}

//------------------------Leetcode Problem 1249------------------------
/*
 * https://leetcode.cn/problems/minimum-remove-to-make-valid-parentheses/
 * 执行用时：12ms 在所有Go提交中击败了57.76%的用户
 * 占用内存：6MB 在所有Go提交中击败了86.34%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", minRemoveToMakeValid(scanner.Text()))
	}
}
