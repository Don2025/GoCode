package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/PzWKhm/
// ------------------------剑指 Offer II Problem 90------------------------
func _rob(nums []int) int {
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 90------------------------
/*
 * https://leetcode.cn/problems/PzWKhm/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了65.73%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", rob(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
