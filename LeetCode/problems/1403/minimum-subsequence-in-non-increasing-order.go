package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order/
//------------------------Leetcode Problem 1403------------------------
func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	var sum, curSum int
	for _, x := range nums {
		sum += x
	}
	var ans []int
	for i := len(nums) - 1; i >= 0; i-- {
		curSum += nums[i]
		sum -= nums[i]
		ans = append(ans, nums[i])
		if curSum > sum {
			break
		}
	}
	return ans
}

//------------------------Leetcode Problem 1403------------------------
/*
 * https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order/
 * 执行用时：8ms 在所有Go提交中击败了17.39%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了60.87%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minSubsequence(nums))
	}
}
