package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums1 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		nums2 := stringArrayToIntArray(strings.Fields(input.Text()))
		println(maxDistance(nums1, nums2))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：128ms 在所有Go提交中击败了75.68%的用户
 * 占用内存：9.4MB 在所有Go提交中击败了18.92%的用户
**/
