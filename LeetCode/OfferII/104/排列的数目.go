package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/D0F0SV/
// ------------------------剑指 Offer II Problem 104------------------------
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if x <= i {
				dp[i] += dp[i-x]
			}
		}
	}
	return dp[target]
}

// ------------------------剑指 Offer II Problem 104------------------------
/*
 * https://leetcode.cn/problems/D0F0SV/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了52.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", combinationSum4(nums, target))
	}
}
