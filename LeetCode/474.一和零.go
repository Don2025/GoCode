package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		m, _ := strconv.Atoi(input.Text())
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		println(findMaxForm(strings.Fields(input.Text()), m, n))
	}
}

/*
 * 执行用时：28ms 在所有Go提交中击败了89.57%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了50.43%的用户
**/
