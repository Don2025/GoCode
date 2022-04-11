package main

import "sort"

func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

/*
 * 执行用时：4ms 在所有Go提交中击败了11.73%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了65.07%的用户
**/
