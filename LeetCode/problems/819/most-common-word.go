package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"unicode"
)

// https://leetcode.cn/problems/most-common-word/
//------------------------Leetcode Problem 819------------------------
func mostCommonWord(paragraph string, banned []string) string {
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	m := make(map[string]int)
	for i := range words {
		m[strings.ToLower(words[i])]++
	}
	ban := make(map[string]bool)
	for _, x := range banned {
		ban[x] = true
	}
	var maxVal int
	var ans string
	for key, val := range m {
		if val > maxVal && !ban[key] {
			maxVal = val
			ans = key
		}
	}
	return ans
}

//------------------------Leetcode Problem 819------------------------
/*
 * // https://leetcode.cn/problems/most-common-word/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3MB 在所有Go提交中击败了30.12%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		paragraph := input.Text()
		input.Scan()
		banned := strings.Fields(input.Text())
		Printf("Output: %v\n", mostCommonWord(paragraph, banned))
	}
}
