package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/NUPfPr/
// ------------------------剑指 Offer II Problem 101------------------------
func canPartition(nums []int) bool {
	var sum int
	for _, x := range nums {
		sum += x
	}
	if sum&1 == 1 {
		return false
	}
	dp := make([]bool, sum>>1+1)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := sum >> 1; j >= 0; j-- {
			if !dp[j] && j >= nums[i-1] {
				dp[j] = dp[j-nums[i-1]]
			}
		}
	}
	return dp[sum>>1]
}

// ------------------------剑指 Offer II Problem 101------------------------
/*
 * https://leetcode.cn/problems/NUPfPr/
 * 执行用时：12ms 在所有Go提交中击败了86.98%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了96.35%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canPartition(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
