package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/IY6buf/
// ------------------------剑指 Offer II Problem 96------------------------
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([]bool, m+1)
	f[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[j] = f[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				f[j] = f[j] || f[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return f[m]
}

// ------------------------剑指 Offer II Problem 96------------------------
/*
 * https://leetcode.cn/problems/IY6buf/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input three strings separated by space on a line:")
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		s1, s2, s3 := arr[0], arr[1], arr[2]
		Printf("Output: %v\n", isInterleave(s1, s2, s3))
		Printf("Input two strings separated by space on a line:")
	}
}
