package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/remove-outermost-parentheses/
//------------------------Leetcode Problem 1021------------------------
func removeOuterParentheses(s string) string {
	var ans []rune
	var level int
	for _, x := range s {
		if x == ')' {
			level--
		}
		if level > 0 {
			ans = append(ans, x)
		}
		if x == '(' {
			level++
		}
	}
	return string(ans)
}

//------------------------Leetcode Problem 1021------------------------
/*
 * https://leetcode.cn/problems/remove-outermost-parentheses/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了64.47%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", removeOuterParentheses(scanner.Text()))
	}
}
