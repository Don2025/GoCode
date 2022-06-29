package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/partition-equal-subset-sum/
//------------------------Leetcode Problem 416------------------------
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

//------------------------Leetcode Problem 416------------------------
/*
 * https://leetcode.cn/problems/partition-equal-subset-sum/
 * 执行用时：21ms 在所有Go提交中击败了93.11%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了71.91%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canPartition(nums))
	}
}
