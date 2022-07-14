package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/backspace-string-compare/
//------------------------Leetcode Problem 844------------------------
func backspaceCompare(s string, t string) bool {
	var skipS, skipT int
	for i, j := len(s)-1, len(t)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
	}
	return true
}

//------------------------Leetcode Problem 844------------------------
/*
 * https://leetcode.cn/problems/backspace-string-compare/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var s, t string
		Sscanf(scanner.Text(), "%s %s", &s, &t)
		Printf("%v\n", backspaceCompare(s, t))
	}
}
