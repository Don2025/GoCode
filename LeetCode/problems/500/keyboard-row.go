package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/keyboard-row/
//------------------------Leetcode Problem 500------------------------
func findWords(words []string) []string {
	keyBoards := []string{"qwertyuiopQWERTYUIOP", "asdfghjklASDFGHJKL", "zxcvbnmZXCVBNM"}
	var ans []string
	for _, word := range words {
		flag1, flag2, flag3 := strings.ContainsAny(word, keyBoards[0]), strings.ContainsAny(word, keyBoards[1]), strings.ContainsAny(word, keyBoards[2])
		if (flag1 && !flag2 && !flag3) || (!flag1 && flag2 && !flag3) || (!flag1 && !flag2 && flag3) {
			ans = append(ans, word)
		}
	}
	return ans
}

//------------------------Leetcode Problem 500------------------------
/*
 * https://leetcode.cn/problems/keyboard-row/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了81.03%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findWords(words))
	}
}
