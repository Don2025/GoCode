package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/rearrange-spaces-between-words/
//------------------------Leetcode Problem 1592------------------------
func reorderSpaces(text string) string {
	space := strings.Count(text, " ")
	words := strings.Fields(text)
	n := len(words) - 1
	if n == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	space, rest := space/n, space%n
	return strings.Join(words, strings.Repeat(" ", space)) + strings.Repeat(" ", rest)
}

//------------------------Leetcode Problem 1592------------------------
/*
 * https://leetcode.cn/problems/rearrange-spaces-between-words/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了25.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", reorderSpaces(scanner.Text()))
	}
}
