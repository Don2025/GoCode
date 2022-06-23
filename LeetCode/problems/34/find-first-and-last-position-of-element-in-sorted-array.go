package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
//------------------------Leetcode Problem 34------------------------
func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}

//------------------------Leetcode Problem 34------------------------
/*
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
 * 执行用时：4ms 在所有Go提交中击败了95.52%的用户
 * 占用内存：3.8MB 在所有Go提交中击败了99.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", searchRange(nums, target))
	}
}
