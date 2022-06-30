package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/ones-and-zeroes/
//------------------------Leetcode Problem 474------------------------
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zeros := strings.Count(strs[i], "0")
		ones := len(strs[i]) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 474------------------------
/*
 * https://leetcode.cn/problems/ones-and-zeroes/
 * 执行用时：28ms 在所有Go提交中击败了89.57%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了50.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		strs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findMaxForm(strs, m, n))
	}
}
