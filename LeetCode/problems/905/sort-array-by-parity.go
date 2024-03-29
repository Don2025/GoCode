package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/sort-array-by-parity/
//------------------------Leetcode Problem 905------------------------
func sortArrayByParity(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]&1 == 1 && nums[right]&1 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		for left < right && nums[left]&1 == 0 {
			left++
		}
		for left < right && nums[right]&1 == 1 {
			right--
		}
	}
	return nums
}

//------------------------Leetcode Problem 905------------------------
/*
 * https://leetcode.cn/problems/sort-array-by-parity/
 * 执行用时：8ms 在所有Go提交中击败了76.16%的用户
 * 占用内存：4.6MB 在所有Go提交中击败了73.51%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := utils.StringToInts(input.Text())
		Printf("Output: %v\n", sortArrayByParity(nums))
	}
}
