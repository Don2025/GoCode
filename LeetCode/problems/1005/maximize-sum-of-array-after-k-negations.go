package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
//------------------------Leetcode Problem 1005------------------------
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = abs(nums[i])
			k--
		}
	}
	if k&1 != 0 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	var ans int
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

//------------------------Leetcode Problem 1005------------------------
/*
 * https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
 * 执行用时：4ms 在所有Go提交中击败了94.59%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了68.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", largestSumAfterKNegations(nums, k))
	}
}
