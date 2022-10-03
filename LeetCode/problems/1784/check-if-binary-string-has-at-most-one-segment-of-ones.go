package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
//------------------------Leetcode Problem 1784------------------------
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

//------------------------Leetcode Problem 1784------------------------
/*
 * https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了28.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %t\n", checkOnesSegment(scanner.Text()))
	}
}
