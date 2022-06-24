package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/decode-ways/
//------------------------Leetcode Problem 91------------------------
func numDecodings(s string) int {
	n := len(s)
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}

//------------------------Leetcode Problem 91------------------------
/*
 * https://leetcode.cn/problems/decode-ways/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了86.38%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", numDecodings(scanner.Text()))
	}
}
