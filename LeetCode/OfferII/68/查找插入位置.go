package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/N6YdxV/
//-------------------------剑指 Offer II Problem 68------------------------
func searchInsert(nums []int, target int) int {
	ans := len(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

//-------------------------剑指 Offer II Problem 68------------------------
/*
 * https://leetcode.cn/problems/N6YdxV/
 * 执行用时：4ms 在所有Go提交中击败了25.65%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了57.62%的用户
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
