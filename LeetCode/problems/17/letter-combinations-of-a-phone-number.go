package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
//------------------------Leetcode Problem 17------------------------
func letterCombinations(digits string) []string {
	var combinations []string
	if len(digits) == 0 {
		return combinations
	}
	phoneKeyboard := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	var backtrack func(string, int)
	backtrack = func(combination string, index int) {
		if index == len(digits) {
			combinations = append(combinations, combination)
		} else {
			t := string(digits[index])
			letters := phoneKeyboard[t]
			for _, x := range letters {
				backtrack(combination+string(x), index+1)
			}
		}
	}
	backtrack("", 0)
	return combinations
}

//------------------------Leetcode Problem 17------------------------
/*
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了61.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", letterCombinations(scanner.Text()))
	}
}
