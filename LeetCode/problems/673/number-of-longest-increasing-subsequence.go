package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
//------------------------Leetcode Problem 673------------------------
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
	}
	var maxCnt int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
			if dp[i] > maxCnt {
				maxCnt = dp[i]
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if maxCnt == dp[i] {
			ans += cnt[i]
		}
	}
	return ans
}

//------------------------Leetcode Problem 673------------------------
/*
 * https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
 * 执行用时：4ms 在所有Go提交中击败了20.90%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了27.07%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findNumberOfLIS(nums))
	}
}
