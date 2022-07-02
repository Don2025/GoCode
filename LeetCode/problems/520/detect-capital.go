package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

// https://leetcode.cn/problems/detect-capital/
//------------------------Leetcode Problem 520------------------------
func detectCapitalUse(word string) bool {
	var cntUpper, cntLower int
	for _, x := range word {
		if unicode.IsLower(x) {
			cntLower++
		} else if unicode.IsUpper(x) {
			cntUpper++
		}
	}
	if cntUpper == len(word) { //全大写
		return true
	}
	if cntLower == len(word) { //全小写
		return true
	}
	if cntUpper == 1 && word[0] < 'a' { //只有首字母大写
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", detectCapitalUse(scanner.Text()))
	}
}

//------------------------Leetcode Problem 520------------------------
/*
 * https://leetcode.cn/problems/detect-capital/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/
