package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/smallest-range-ii/
//------------------------Leetcode Problem 910------------------------
func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		high := max(nums[n-1]-k, nums[i]+k)
		low := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, high-low)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 910------------------------
/*
 * https://leetcode.cn/problems/smallest-range-ii/
 * 执行用时：20ms 在所有Go提交中击败了77.97%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了83.05%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", smallestRangeII(nums, k))
	}
}
