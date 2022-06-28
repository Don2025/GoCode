package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/reverse-string/
//------------------------Leetcode Problem 344------------------------
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

//------------------------Leetcode Problem 344------------------------
/*
 * https://leetcode.cn/problems/reverse-string/
 * 执行用时：28ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了78.62%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := []byte(scanner.Text())
		reverseString(s)
		Printf("Output: %s\n", s)
	}
}
