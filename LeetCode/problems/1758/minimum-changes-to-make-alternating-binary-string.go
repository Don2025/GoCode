package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://leetcode.cn/problems/minimum-changes-to-make-alternating-binary-string/
//------------------------Leetcode Problem 1758------------------------
func minOperations(s string) int {
	cnt := 0
	for i, c := range s {
		if int(c-'0') != i%2 {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1758------------------------
/*
 * https://leetcode.cn/problems/minimum-changes-to-make-alternating-binary-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了67.52%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("Output: %v\n", minOperations(scanner.Text()))
	}
}
