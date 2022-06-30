package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/
//------------------------Leetcode Problem 442------------------------
func findDuplicates(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	var ans []int
	for i := range nums {
		if nums[i]-1 != i {
			ans = append(ans, nums[i])
		}
	}
	return ans
}

//------------------------Leetcode Problem 442------------------------
/*
 * https://leetcode.cn/problems/find-all-duplicates-in-an-array/
 * 执行用时：44ms 在所有Go提交中击败了50.88%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了63.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findDuplicates(nums))
	}
}
