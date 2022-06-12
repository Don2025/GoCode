package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/find-and-replace-pattern/
//------------------------Leetcode Problem 890------------------------
func findAndReplacePattern(words []string, pattern string) (ans []string) {
	match := func(word, pattern string) bool {
		m := make(map[rune]byte)
		for i, x := range word {
			y := pattern[i]
			if m[x] == 0 {
				m[x] = y
			} else if m[x] != y {
				return false
			}
		}
		return true
	}
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			ans = append(ans, word)
		}
	}
	return
}

//------------------------Leetcode Problem 890------------------------
/*
 * https://leetcode.cn/problems/find-and-replace-pattern/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了32.14%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of strings separated by space:")
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Input a line of string:")
		scanner.Scan()
		pattern := scanner.Text()
		Printf("Output: %v\n", findAndReplacePattern(words, pattern))
		Printf("Input a line of strings separated by space:")
	}
}
