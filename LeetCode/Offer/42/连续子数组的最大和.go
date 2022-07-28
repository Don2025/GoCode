package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

/* 动态规划求解
 * 执行用时：16ms 在所有Go提交中击败了99.68%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了17.32%的用户

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}
**/
// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// ------------------------剑指 Offer I Problem 42------------------------
func maxSubArray(nums []int) int {
	var sum int
	ans := math.MinInt32
	for _, x := range nums {
		if sum <= 0 {
			sum = x
		} else {
			sum += x
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer I Problem 42------------------------
/* 前缀和
 * https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
 * 执行用时：20ms 在所有Go提交中击败了97.62%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了98.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxSubArray(nums))
	}
}
