package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func stoneGame(piles []int) bool {
	dp := make([][]int, len(piles))
	// 只有一堆石头时，先手能领先的最大分值就是这堆石头的数量
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]int, len(piles))
		dp[i][i] = piles[i]
	}
	for i := len(piles) - 2; i >= 0; i-- {
		for j := i + 1; j < len(piles); j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][len(piles)-1] > 0
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
		println(stoneGame(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：8ms 在所有Go提交中击败了32.03%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了45.75%的用户
**/
