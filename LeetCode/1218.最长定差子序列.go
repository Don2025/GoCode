package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func longestSubsequence(arr []int, difference int) int {
	var ans int
	dp := map[int]int{}
	for _, x := range arr {
		dp[x] = dp[x-difference] + 1
		ans = max(ans, dp[x])
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
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		difference, _ := strconv.Atoi(input.Text())
		println(longestSubsequence(arr, difference))
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
 * 执行用时：96ms 在所有Go提交中击败了82.61%的用户
 * 占用内存：9.1MB 在所有Go提交中击败了47.83%的用户
**/
