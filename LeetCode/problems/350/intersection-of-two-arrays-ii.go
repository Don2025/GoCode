package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
//------------------------Leetcode Problem 350------------------------
func intersect(nums1 []int, nums2 []int) []int {
	cnt := make(map[int]int)
	var ans []int
	for _, x := range nums1 {
		cnt[x]++
	}
	for _, x := range nums2 {
		if cnt[x] > 0 {
			ans = append(ans, x)
			cnt[x]--
		}
	}
	return ans
}

//------------------------Leetcode Problem 350------------------------
/*
 * https://leetcode.cn/problems/intersection-of-two-arrays-ii/
 * 执行用时：4ms 在所有Go提交中击败了73.88%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了36.79%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("%v\n", intersect(nums1, nums2))
	}
}
