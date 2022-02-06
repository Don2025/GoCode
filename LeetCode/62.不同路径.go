package main

import (
	"bufio"
	"os"
	"strconv"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		m, _ := strconv.Atoi(input.Text())
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		println(uniquePaths(m, n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.19%的用户
**/
