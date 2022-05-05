package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxValue(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n+1)
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i-1][j-1]
		}
	}
	return dp[n]
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
		n, _ := strconv.Atoi(input.Text())
		grid := make([][]int, n)
		for i := range grid {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(maxValue(grid))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：4ms 在所有Go提交中击败了97.96%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了61.64%的用户
**/
