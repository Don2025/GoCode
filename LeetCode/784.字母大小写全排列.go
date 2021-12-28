package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := letterCasePermutation(input.Text())
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了73.75%的用户
**/
