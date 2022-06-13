package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/cyJERH/
// ------------------------剑指 Offer II Problem 92------------------------
func minFlipsMonoIncr(s string) int {
	var dp0, dp1 int
	for _, ch := range s {
		dp0New, dp1New := dp0, min(dp0, dp1)
		if ch == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return min(dp0, dp1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 92------------------------
/*
 * https://leetcode.cn/problems/cyJERH/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string containing only 0 and 1:")
	for scanner.Scan() {
		str := scanner.Text()
		Printf("Output: %v\n", minFlipsMonoIncr(str))
		Printf("Input a line of string containing only 0 and 1:")
	}
}
