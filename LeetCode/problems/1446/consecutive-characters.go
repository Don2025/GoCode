package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/consecutive-characters/
//------------------------Leetcode Problem 1446------------------------
func maxPower(s string) int {
	cnt, ans := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1446------------------------
/*
 * https://leetcode.cn/problems/consecutive-characters/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了58.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", maxPower(scanner.Text()))
	}
}
