package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	//偷第一间房
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	ans := dp[len(dp)-2]
	//偷最后一间房
	dp[0], dp[1] = 0, 0
	for i := 2; i < len(nums)+1; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	ans = max(ans, dp[len(dp)-1])
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
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(rob(nums))
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
 * 占用内存：2MB 在所有Go提交中击败了47.22%的用户
**/
