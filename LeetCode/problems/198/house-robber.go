package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/house-robber/
//------------------------Leetcode Problem 198------------------------
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 198------------------------
/*
 * https://leetcode.cn/problems/house-robber/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了61.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", rob(nums))
	}
}
