package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				dp[i] = dp[i] + dp[i-coin]
			}
		}
	}
	return dp[amount]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		coins := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		amount, _ := strconv.Atoi(input.Text())
		println(change(amount, coins))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了62.00%的用户
**/
