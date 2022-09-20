package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
//------------------------Leetcode Problem 698------------------------
func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, x := range nums {
		all += x
	}
	if all%k != 0 {
		return false
	}
	target := all / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > target {
		return false
	}
	dp := make([]bool, 1<<n)
	dp[0] = true
	sum := make([]int, 1<<n)
	for i, x := range dp {
		if !x {
			continue
		}
		for j, num := range nums {
			if sum[i]+num > target {
				break
			}
			if i>>j&1 == 0 {
				next := i | 1<<j
				if !dp[next] {
					sum[next] = (sum[i] + num) % target
					dp[next] = true
				}
			}
		}
	}
	return dp[(1<<n)-1]
}

// ------------------------Leetcode Problem 698------------------------
/*
 * https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
 * 执行用时：44ms 在所有Go提交中击败了55.87%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了7.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canPartitionKSubsets(nums))
	}
}
