package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/
//------------------------Leetcode Problem 801------------------------
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	a, b := 0, 1
	for i := 1; i < n; i++ {
		at, bt := n, n
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			at, bt = a, b+1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			at, bt = min(at, b), min(bt, a+1)
		}
		a, b = at, bt
	}
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 801------------------------
/*
 * https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/
 * 执行用时：160ms 在所有Go提交中击败了13.64%的用户
 * 占用内存：10MB 在所有Go提交中击败了45.45%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minSwap(nums1, nums2))
	}
}
