package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/longest-increasing-subsequence/
//------------------------Leetcode Problem 300------------------------
func lengthOfLIS(nums []int) int {
	var dp []int
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		}
	}
	return len(dp)
}

//------------------------Leetcode Problem 300------------------------
/*
 * https://leetcode.cn/problems/longest-increasing-subsequence/
 * 执行用时：4ms 在所有Go提交中击败了98.99%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了87.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", lengthOfLIS(nums))
	}
}
