package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/determine-if-string-halves-are-alike/
//------------------------Leetcode Problem 1704------------------------
func halvesAreAlike(s string) bool {
	var cnt, i int
	for i < len(s)/2 {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			cnt++
		}
		i++
	}
	for i < len(s) {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			cnt--
		}
		i++
	}
	return cnt == 0
}

//------------------------Leetcode Problem 1704------------------------
/*
 * https://leetcode.cn/problems/determine-if-string-halves-are-alike/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了79.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", halvesAreAlike(scanner.Text()))
	}
}
