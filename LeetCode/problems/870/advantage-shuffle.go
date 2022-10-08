package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/advantage-shuffle/
//------------------------Leetcode Problem 870------------------------
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	sort.Ints(nums1)
	ans := make([]int, n)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})
	l, r := 0, n-1
	for _, x := range nums1 {
		if x > nums2[idx[l]] {
			ans[idx[l]] = x
			l++
		} else {
			ans[idx[r]] = x
			r--
		}
	}
	return ans
}

//------------------------Leetcode Problem 870------------------------
/*
 * https://leetcode.cn/problems/advantage-shuffle/
 * 执行用时：128ms 在所有Go提交中击败了88.00%的用户
 * 占用内存：9.9MB 在所有Go提交中击败了83.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", advantageCount(nums1, nums2))
	}
}
