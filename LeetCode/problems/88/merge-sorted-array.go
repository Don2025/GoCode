package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/merge-sorted-array/
//------------------------Leetcode Problem 88------------------------
func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := len(nums1) - 1; m >= 0 && n >= 0; i-- {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
	if n >= 0 {
		copy(nums1[0:n+1], nums2[0:n+1])
	}
}

//------------------------Leetcode Problem 88------------------------
/*
 * https://leetcode.cn/problems/merge-sorted-array/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了74.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		merge(nums1, m, nums2, n)
		Printf("Output: %v\n", nums1)
	}
}
