package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

//  https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/
//------------------------Leetcode Problem 1662------------------------
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var p1, p2, i, j int
	for p1 < len(word1) && p2 < len(word2) {
		if word1[p1][i] != word2[p2][j] {
			return false
		}
		i++
		if i == len(word1[p1]) {
			p1++
			i = 0
		}
		j++
		if j == len(word2[p2]) {
			p2++
			j = 0
		}
	}
	return p1 == len(word1) && p2 == len(word2)
}

//------------------------Leetcode Problem 1662------------------------
/*
 * https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了80.95%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word1 := strings.Fields(scanner.Text())
		scanner.Scan()
		word2 := strings.Fields(scanner.Text())
		Printf("Output: %t\n", arrayStringsAreEqual(word1, word2))
	}
}
