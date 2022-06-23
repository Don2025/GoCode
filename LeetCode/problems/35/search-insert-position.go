package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/search-insert-position/
//------------------------Leetcode Problem 35------------------------
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else { // target > nums[mid]
			l = mid + 1
		}
	}
	return l
}

//------------------------Leetcode Problem 35------------------------
/*
 * https://leetcode.cn/problems/search-insert-position/
 * 执行用时：4ms 在所有Go提交中击败了71.71%的用户
 * 占用内存：3MB 在所有Go提交中击败了49.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", searchInsert(nums, target))
	}
}
