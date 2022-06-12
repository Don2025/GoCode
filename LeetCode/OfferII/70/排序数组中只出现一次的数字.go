package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/skFtm2/
// ------------------------剑指 Offer II Problem 70------------------------
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// ------------------------剑指 Offer II Problem 70------------------------
/*
 * https://leetcode.cn/problems/skFtm2/
 * 执行用时：4ms 在所有Go提交中击败了96.40%的用户
 * 占用内存：3.8MB 在所有Go提交中击败了60.79%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNonDuplicate(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
