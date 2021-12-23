package main

import (
	"bufio"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var ans int
	for l, r := 0, 0; r < len(s); r++ {
		l = max(l, m[s[r]])
		m[s[r]] = r + 1
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(lengthOfLongestSubstring(input.Text()))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了45.79%的用户
 * 占用内存：3MB 在所有Go提交中击败了44.91%的用户
**/
