package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/YaVDxD/
// ------------------------剑指 Offer II Problem 102------------------------
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, x := range nums {
		sum += x
	}
	diff := sum - target
	if diff < 0 || diff&1 == 1 {
		return 0
	}
	neg := diff >> 1
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

// ------------------------剑指 Offer II Problem 102------------------------
/*
 * https://leetcode.cn/problems/YaVDxD/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了62.37%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findTargetSumWays(nums, target))
		Printf("Input a line of numbers separated by space:")
	}
}
