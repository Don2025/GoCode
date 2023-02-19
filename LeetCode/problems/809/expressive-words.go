package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/expressive-words/
//------------------------Leetcode Problem 809------------------------
func expressiveWords(s string, words []string) (cnt int) {
	isStretchy := func(s, t string) bool {
		i, j, m, n := 0, 0, len(s), len(t)
		for i < m && j < n {
			if s[i] != t[j] {
				return false
			}
			ch := s[i]
			var cntI, cntJ int
			for i < m && ch == s[i] {
				i++
				cntI++
			}
			for j < n && ch == t[j] {
				j++
				cntJ++
			}
			if cntI < cntJ || cntI < 3 && cntI > cntJ {
				return false
			}
		}
		return i == m && j == n
	}
	for _, word := range words {
		if isStretchy(s, word) {
			cnt++
		}
	}
	return
}

//------------------------Leetcode Problem 809------------------------
/*
 * https://leetcode.cn/problems/expressive-words/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了63.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", expressiveWords(s, words))
	}
}
