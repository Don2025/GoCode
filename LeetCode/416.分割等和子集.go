package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func canPartition(nums []int) bool {
	var sum int
	for _, x := range nums {
		sum += x
	}
	if sum&1 == 1 {
		return false
	}
	t := sum >> 1
	dp := make([]int, t+1)
	for _, x := range nums {
		for j := t; j >= x; j-- {
			dp[j] = max(dp[j], dp[j-x]+x)
		}
	}
	if dp[t] == t {
		return true
	}
	return false
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
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(canPartition(nums))
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
 * 执行用时：21ms 在所有Go提交中击败了93.11%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了71.91%的用户
**/
