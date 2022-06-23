package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
//------------------------Leetcode Problem 4------------------------
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr []float64
	for _, x := range nums1 {
		arr = append(arr, float64(x))
	}
	for _, x := range nums2 {
		arr = append(arr, float64(x))
	}
	sort.Float64s(arr)
	n := len(arr) >> 1
	if len(arr)&1 == 0 {
		return (arr[n] + arr[n-1]) / 2
	} else {
		return arr[n]
	}
}

//------------------------Leetcode Problem 4------------------------
/*
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/
 * 执行用时：16ms 在所有Go提交中击败了48.03%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了11.68%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		nums2 := utils.StringToInts(scanner.Text())
		Printf("%v\n", findMedianSortedArrays(nums1, nums2))
	}
}
