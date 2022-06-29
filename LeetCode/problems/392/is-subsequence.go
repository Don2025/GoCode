package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/is-subsequence/
//------------------------Leetcode Problem 392------------------------
func isSubsequence(s string, t string) bool {
	var i int
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

//------------------------Leetcode Problem 392------------------------
/*
 * https://leetcode.cn/problems/is-subsequence/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了89.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", isSubsequence(arr[0], arr[1]))
	}
}
