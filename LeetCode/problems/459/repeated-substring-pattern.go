package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/repeated-substring-pattern/
//------------------------Leetcode Problem 459------------------------
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

//------------------------Leetcode Problem 459------------------------
/*
 * https://leetcode.cn/problems/repeated-substring-pattern/
 * 执行用时：4ms 在所有Go提交中击败了90.24%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了45.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", repeatedSubstringPattern(scanner.Text()))
	}
}
