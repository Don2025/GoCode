package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3]int, n) //dp[i][x]第i天进入(处于)x状态（0.不持股 1.持股 2.冷冻期）
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[n-1][0], dp[n-1][2])
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
		prices := stringArrayToIntArray(strings.Fields(input.Text()))
		println(maxProfit(prices))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了61.14%的用户
**/
