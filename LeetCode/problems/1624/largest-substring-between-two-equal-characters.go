package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
//------------------------Leetcode Problem 1624------------------------
func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	m := [26]int{}
	for i := range m {
		m[i] = -1
	}
	for i, x := range s {
		x -= 'a'
		if m[x] == -1 {
			m[x] = i
		} else {
			ans = max(ans, i-m[x]-1)
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

//------------------------Leetcode Problem 1624------------------------
/*
 * https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了56.52%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %d\n", maxLengthBetweenEqualCharacters(scanner.Text()))
	}
}
