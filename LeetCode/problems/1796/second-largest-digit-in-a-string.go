package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

// https://leetcode.cn/problems/second-largest-digit-in-a-string/
//------------------------Leetcode Problem 1796------------------------
func secondHighest(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if unicode.IsDigit(c) {
			n := int(c - '0')
			if n > first {
				first, second = n, first
			} else if n < first && n > second {
				second = n
			}
		}
	}
	return second
}

//------------------------Leetcode Problem 1796------------------------
/*
 * https://leetcode.cn/problems/second-largest-digit-in-a-string/
 * 执行用时：0ms 在所有Go提交中击败了30.30%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了72.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", secondHighest(scanner.Text()))
	}
}
