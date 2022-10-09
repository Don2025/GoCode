package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/score-of-parentheses/
//------------------------Leetcode Problem 856------------------------
func scoreOfParentheses(s string) int {
	var depth, ans int
	for i, x := range s {
		if x == '(' {
			depth++
		} else {
			depth--
			if s[i-1] == '(' {
				ans += 1 << depth
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 856------------------------
/*
 * https://leetcode.cn/problems/score-of-parentheses/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了46.34%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", scoreOfParentheses(scanner.Text()))
	}
}
