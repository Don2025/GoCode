package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/reverse-words-in-a-string-iii/
//------------------------Leetcode Problem 557------------------------
func reverseWords(s string) string {
	var i int
	ans := []byte(s)
	for i < len(s) {
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		l, r := start, i-1
		for l < r {
			ans[l], ans[r] = s[r], s[l]
			l++
			r--
		}
		for i < len(s) && s[i] == ' ' {
			i++
		}
	}
	return string(ans)
}

//------------------------Leetcode Problem 557------------------------
/*
 * https://leetcode.cn/problems/reverse-words-in-a-string-iii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了80.51%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %s\n", reverseWords(scanner.Text()))
	}
}
