package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/valid-parentheses/
//------------------------Leetcode Problem 20------------------------
func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte
	for i := range s {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

//------------------------Leetcode Problem 20------------------------
/*
 * https://leetcode.cn/problems/valid-parentheses/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了93.59%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", isValid(scanner.Text()))
	}
}
