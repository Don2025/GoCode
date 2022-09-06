package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/
//------------------------Leetcode Problem 828------------------------
func uniqueLetterString(s string) int {
	var ans int
	for ch := 'A'; ch <= 'Z'; ch++ {
		for i, l, r := 0, -1, -1; i < len(s); i++ {
			if s[i] == byte(ch) {
				l, r = r, i
			}
			ans += r - l
		}
	}
	return ans % 1000000007
}

//------------------------Leetcode Problem 828------------------------
/*
 * https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/
 * 执行用时：24ms 在所有Go提交中击败了42.86%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了60.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", uniqueLetterString(scanner.Text()))
	}
}
