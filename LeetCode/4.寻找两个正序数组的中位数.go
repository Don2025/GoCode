package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums1 := utils.StringArrayToIntArray(strings.Fields(input.Text()))
		nums2 := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findMedianSortedArrays(nums1, nums2))
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了48.03%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了11.68%的用户
**/
