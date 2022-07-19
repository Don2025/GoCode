package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-sum-circular-subarray/
//------------------------Leetcode Problem 918------------------------
func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, curMax, curMin := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		curMax = max(curMax+nums[i], nums[i])
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+nums[i], nums[i])
		minSum = min(minSum, curMin)
	}
	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 918------------------------
/*
 * https://leetcode.cn/problems/maximum-sum-circular-subarray/
 * 执行用时：56ms 在所有Go提交中击败了80.51%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了90.68%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxSubarraySumCircular(nums))
	}
}
