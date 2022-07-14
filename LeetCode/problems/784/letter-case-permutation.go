package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

// https://leetcode.cn/problems/letter-case-permutation/
//------------------------Leetcode Problem 784------------------------
func letterCasePermutation(s string) []string {
	var ans []string
	var backtrack func([]rune, int)
	backtrack = func(str []rune, i int) {
		if i == len(str) {
			ans = append(ans, string(str))
			return
		}
		if unicode.IsLetter(str[i]) {
			backtrack(str, i+1)
			str[i] ^= 32 // 大写变小写 小写变大写
			backtrack(str, i+1)
		} else {
			backtrack(str, i+1)
		}
	}
	backtrack([]rune(s), 0)
	return ans
}

//------------------------Leetcode Problem 784------------------------
/*
 * https://leetcode.cn/problems/letter-case-permutation/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了73.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", letterCasePermutation(scanner.Text()))
	}
}
