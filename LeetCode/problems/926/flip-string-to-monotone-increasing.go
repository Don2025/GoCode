package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
//------------------------Leetcode Problem 926------------------------
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
	if a > b {
		return b
	}
	return a
}

//------------------------Leetcode Problem 926------------------------
/*
 * https://leetcode.cn/problems/flip-string-to-monotone-increasing/
 * 执行用时：16ms 在所有Go提交中击败了17.53%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了42.93%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string:")
	for scanner.Scan() {
		Printf("Output: %v\n", minFlipsMonoIncr(scanner.Text()))
		Printf("Input a line of string:")
	}
}
