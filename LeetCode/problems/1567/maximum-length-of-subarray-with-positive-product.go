package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product/
//------------------------Leetcode Problem 1567------------------------
func getMaxLen(nums []int) int {
	var pos, neg int
	if nums[0] > 0 {
		pos = 1
	}
	if nums[0] < 0 {
		neg = 1
	}
	ans := pos
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			pos++
			if neg > 0 {
				neg++
			} else {
				neg = 0
			}
		} else if nums[i] < 0 {
			t := pos
			if neg > 0 {
				pos = neg + 1
			} else {
				pos = 0
			}
			neg = t + 1
		} else {
			pos, neg = 0, 0
		}
		ans = max(ans, pos)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1567------------------------
/*
 * https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product/
 * 执行用时：100ms 在所有Go提交中击败了22.06%的用户
 * 占用内存：9.9MB 在所有Go提交中击败了84.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", getMaxLen(nums))
	}
}
