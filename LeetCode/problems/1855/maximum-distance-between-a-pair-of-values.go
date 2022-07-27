package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-distance-between-a-pair-of-values/
//------------------------Leetcode Problem 1855------------------------
func maxDistance(nums1 []int, nums2 []int) int {
	var ans, i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			ans = max(ans, j-i)
			j++
		} else {
			i++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1855------------------------
/*
 * https://leetcode.cn/problems/maximum-distance-between-a-pair-of-values/
 * 执行用时：128ms 在所有Go提交中击败了75.68%的用户
 * 占用内存：9.4MB 在所有Go提交中击败了18.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", maxDistance(nums1, nums2))
	}
}
