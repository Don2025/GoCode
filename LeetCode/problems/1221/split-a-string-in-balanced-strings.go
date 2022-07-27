package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/split-a-string-in-balanced-strings/
//------------------------Leetcode Problem 1221------------------------
func balancedStringSplit(s string) int {
	var cnt, n int
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			n++
		} else if s[i] == 'L' {
			n--
		}
		if n == 0 {
			cnt++
		}
	}
	return cnt
}

//------------------------Leetcode Problem 1221------------------------
/*
 * https://leetcode.cn/problems/split-a-string-in-balanced-strings/
 * 执行用时：0ms 在所有Go提交中击败了11.62%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了77.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", balancedStringSplit(scanner.Text()))
	}
}
