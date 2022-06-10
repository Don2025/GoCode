package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func coinChange(coins []int, amount int) int {
	maxVal := amount + 1
	dp := make([]int, maxVal)
	for i := 1; i < len(dp); i++ {
		dp[i] = maxVal
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		coins := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		amount, _ := strconv.Atoi(input.Text())
		println(coinChange(coins, amount))
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
 * 执行用时：8ms 在所有Go提交中击败了93.02%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了58.53%的用户
**/
