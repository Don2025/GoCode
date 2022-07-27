package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/longest-happy-prefix/
//------------------------Leetcode Problem 1392------------------------
func longestPrefix(s string) string {
	n := len(s)
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
	return s[0:next[n-1]]
}

//------------------------Leetcode Problem 1392------------------------
/*
 * https://leetcode.cn/problems/longest-happy-prefix/
 * 执行用时：12ms 在所有Go提交中击败了85.94%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了37.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", longestPrefix(scanner.Text()))
	}
}
