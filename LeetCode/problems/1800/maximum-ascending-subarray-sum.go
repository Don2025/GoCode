package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-ascending-subarray-sum/
// ------------------------Leetcode Problem 1800------------------------
func maxAscendingSum(nums []int) int {
	var ans, sum int
	for i := range nums {
		if i == 0 || nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
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

//------------------------Leetcode Problem 1800------------------------
/*
 * https://leetcode.cn/problems/maximum-ascending-subarray-sum/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了98.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", maxAscendingSum(nums))
	}
}
